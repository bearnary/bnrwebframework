package bnrwebframework

import (
	"github.com/gin-gonic/gin"
	redocmiddleware "github.com/go-openapi/runtime/middleware"
)

func (c *defaultGin) RedocInit(route string, filePath string) {
	opts := redocmiddleware.RedocOpts{SpecURL: filePath}
	sh := redocmiddleware.Redoc(opts, nil)

	c.server.GET(route, func(ctx *gin.Context) {
		sh.ServeHTTP(ctx.Writer, ctx.Request)
	})
	c.server.StaticFile(filePath, filePath)
}
