package bnrwebframework

import (
	"net/http"
)

func (c *defaultJWT) ValidateTokenMiddleware(ctx *Context) {
	if _, err := c.ParseAccessToken(ctx); err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
