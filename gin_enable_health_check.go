package bnrwebframework

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *defaultGin) EnableHealthCheck(path string, healthCheckExtensionFunc func() error) {
	c.server.GET(path, func(c *gin.Context) {

		if healthCheckExtensionFunc != nil {
			err := healthCheckExtensionFunc()
			if err != nil {
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "health  success",
		})
	})
}
