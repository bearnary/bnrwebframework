package bnrwebframework

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

var signingMethod = jwt.SigningMethodHS256

type JWTConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	TimeoutMinute      int
	MaxRefreshMinute   int
}

type Token struct {
	Token     string
	ExpiredAt time.Time
}

type defaultJWT struct {
	cfg *JWTConfig
}

func NewJWT(cfg *JWTConfig) JWT {
	return &defaultJWT{
		cfg: cfg,
	}
}

func (c *defaultJWT) GenerateAccessToken(model interface{}) (*Token, error) {
	return generateToken(model, c.cfg.AccessTokenSecret, c.cfg.TimeoutMinute)
}

func (c *defaultJWT) GenerateRefreshToken(model interface{}) (*Token, error) {
	return generateToken(model, c.cfg.RefreshTokenSecret, c.cfg.MaxRefreshMinute)
}

func generateToken(model interface{}, secret string, timeoutMinute int) (*Token, error) {
	claims := jwt.MapClaims{}
	now := time.Now()
	expirationTime := now.Add(time.Minute * time.Duration(timeoutMinute))
	claims["exp"] = expirationTime.Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem()
	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag.Get("jwt")
		name := t.Field(i).Name
		value := v.FieldByName(name).Interface()
		claims[tag] = value
	}
	token, err := jwt.NewWithClaims(signingMethod, claims).SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &Token{
		Token:     token,
		ExpiredAt: expirationTime,
	}, nil
}

func (c *defaultJWT) ParseAccessToken(ctx *Context) (map[string]interface{}, error) {
	token := c.ExtractToken(ctx)
	return c.parseToken(token, c.cfg.AccessTokenSecret)
}

func (c *defaultJWT) ParseRefreshToken(token string) (map[string]interface{}, error) {
	return c.parseToken(token, c.cfg.RefreshTokenSecret)
}

func (c *defaultJWT) parseToken(token string, secret string) (map[string]interface{}, error) {
	secretKeyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
	t, err := jwt.Parse(token, secretKeyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func (c *defaultJWT) ExtractToken(ctx *Context) string {
	// Extract the token string from the context
	value := ctx.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(value, "Bearer ")
	return tokenString
}
