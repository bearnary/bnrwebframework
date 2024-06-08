package bnrwebframework

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const pathKey = "path"

type CustomBinder struct{}

// BindAndValidate bind and validate form input then return error in human-readable strings
func (c *Context) BindAndValidate(i interface{}) error {

	err := c.ShouldBind(i)
	if err != nil {
		eStrs := parseError(err)
		eStr := strings.Join(eStrs, ",")
		return errors.New(eStr)
	}

	c.parsePathParams(i)

	c.Parameters = i
	return nil
}

// BindJSONAndValidate bind and validate json input then return error in human-readable strings
func (c *Context) BindJSONAndValidate(i interface{}) error {

	err := c.BindJSON(i)
	if err != nil {
		eStrs := parseError(err)
		eStr := strings.Join(eStrs, ",")
		return errors.New(eStr)
	}

	c.parsePathParams(i)

	c.Parameters = i
	return nil
}

// parsePathParams parse path parameters and store in context
func (c *Context) parsePathParams(form interface{}) {
	formValue := reflect.ValueOf(form)
	if formValue.Kind() == reflect.Ptr {
		formValue = formValue.Elem()
	}
	t := reflect.TypeOf(formValue.Interface())
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(pathKey)
		if tag != "" {
			fieldName := t.Field(i).Name
			paramValue := formValue.FieldByName(fieldName)
			if paramValue.IsValid() {
				paramValue.Set(reflect.ValueOf(c.Param(tag)))
			}
		}
	}
}

// parseError takes an error or multiple errors and attempts to determine the best path to convert them into
// human-readable strings
func parseError(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := err.(type) {
		case validator.ValidationErrors:
			// if the type is validator.ValidationErrors then it's actually an array of validator.FieldError so we'll
			// loop through each of those and convert them one by one
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			// similarly, if the error is an unmarshalling error we'll parse it into another, more readable string format
			out = append(out, parseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}
	return out
}

// parseFieldError parse field error to human-readable format
func parseFieldError(e validator.FieldError) string {
	// workaround to the fact that the `gt|gtfield=Start` gets passed as an entire tag for some reason
	// https://github.com/go-playground/validator/issues/926

	fieldPrefix := fmt.Sprintf("The field %s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)
	default:
		// if it's a tag for which we don't have a good format string yet we'll try using the default english translator
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		} else {
			return fmt.Errorf("%v", e).Error()
		}
	}
}

// parseMarshallingError parse marshalling error to human-readable string
func parseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("The field %s must be a %s", e.Field, e.Type.String())
}
