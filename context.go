package bnrwebframework

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	Parameters interface{}
}

// NewContext New create new context from gin's context
func NewContext(ctx *gin.Context) *Context {
	return &Context{
		Context: ctx,
	}
}
