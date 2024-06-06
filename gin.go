package bnrwebframework

import "github.com/gin-gonic/gin"

type Gin interface {
	Start(port string) error

	Group(relativePath string) RouterGroup
	RegisterAPIVersionGroup() RouterGroup
	APIVersionGroup() RouterGroup

	Engine() *gin.Engine

	Use(middlewares ...HandlerFunc)
	EnableDefaultLogMiddleware()
	EnableHealthCheck(path string, healthCheckExtensionFunc func() error)

	RedocInit(route string, filepath string)
}
