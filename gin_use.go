package bnrwebframework

import "github.com/gin-gonic/gin"

func (c *defaultGin) Use(middlewares ...HandlerFunc) {

	var handlers []gin.HandlerFunc
	for _, h := range middlewares {
		handler := newGinHandler(h)
		handlers = append(handlers, handler)
	}
	c.server.Use(handlers...)
}
