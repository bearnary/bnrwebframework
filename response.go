package bnrwebframework

import (
	"net/http"

	"github.com/oneononex/oolib/ooerrors"
)

const (
	ErrHTTPCode = 999
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// wrapResponse create default response with code and message
func (c *Context) wrapResponse(v interface{}) BaseResponse {
	resp := BaseResponse{
		Code:    0,
		Message: "",
		Data:    v,
	}
	return resp
}

// CreateResponseSuccess create default response from struct
func (c *Context) CreateResponseSuccess(v interface{}, status ...int) {
	resp := c.wrapResponse(v)

	s := http.StatusOK
	if len(status) > 0 {
		s = status[0]
	}

	c.JSON(s, resp)
}

func (c *Context) CreateResponseError(code int, message string, status ...int) {
	resp := BaseResponse{
		Code:    code,
		Message: message,
	}

	s := ErrHTTPCode
	if len(status) > 0 {
		s = status[0]
	}

	c.JSON(s, resp)
}

func (c *Context) CreateResponseErrorWithOOError(vErr ooerrors.Error, status ...int) {
	resp := BaseResponse{
		Code:    vErr.Code(),
		Message: vErr.Message(),
	}

	s := ErrHTTPCode
	if len(status) > 0 {
		s = status[0]
	}

	c.JSON(s, resp)
}
