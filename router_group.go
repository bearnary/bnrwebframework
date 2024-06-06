package bnrwebframework

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup interface {
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup

	POST(relativePath string, handlers ...HandlerFunc)
	GET(relativePath string, handlers ...HandlerFunc)
	PUT(relativePath string, handlers ...HandlerFunc)
	PATCH(relativePath string, handlers ...HandlerFunc)
	DELETE(relativePath string, handlers ...HandlerFunc)

	USE(middlewares ...HandlerFunc)
}

type defaultRouterGroup struct {
	g gin.RouterGroup
}

func convertGinHandlers(hs ...HandlerFunc) []gin.HandlerFunc {

	var handlers []gin.HandlerFunc
	for _, h := range hs {
		handler := newGinHandler(h)
		handlers = append(handlers, handler)
	}
	return handlers
}

func newGinHandler(h HandlerFunc) gin.HandlerFunc {
	handler := func(context *gin.Context) {
		if !context.IsAborted() {
			h(NewContext(context))
		}
	}
	return handler
}

// Group create new sub RouterGroup from relativePath
func (rg *defaultRouterGroup) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {

	hs := convertGinHandlers(handlers...)

	g := rg.g.Group(relativePath, hs...)
	return &defaultRouterGroup{
		g: *g,
	}
}

// POST create POST request with handler
func (rg *defaultRouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	hs := convertGinHandlers(handlers...)
	rg.g.POST(relativePath, hs...)
}

// GET create GET request with handler
func (rg *defaultRouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	hs := convertGinHandlers(handlers...)
	rg.g.GET(relativePath, hs...)
}

// PUT create PUT request with handler
func (rg *defaultRouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
	hs := convertGinHandlers(handlers...)
	rg.g.PUT(relativePath, hs...)
}

// PATCH create PATCH request with handler
func (rg *defaultRouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) {
	hs := convertGinHandlers(handlers...)
	rg.g.PATCH(relativePath, hs...)
}

// DELETE create DELETE request with handler
func (rg *defaultRouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
	hs := convertGinHandlers(handlers...)
	rg.g.DELETE(relativePath, hs...)
}

// USE create USE request with handler
func (rg *defaultRouterGroup) USE(handlers ...HandlerFunc) {
	hs := convertGinHandlers(handlers...)
	rg.g.Use(hs...)
}
