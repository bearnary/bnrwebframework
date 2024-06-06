package bnrwebframework

import "github.com/gin-gonic/gin"

func (c *defaultGin) EnableDefaultLogMiddleware() {
	c.server.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/"))
}
