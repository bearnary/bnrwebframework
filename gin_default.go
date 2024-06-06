package bnrwebframework

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type defaultGin struct {
	server          *gin.Engine
	apiVersionGroup RouterGroup
}

// New create new Gin interface
func New(config *Config) Gin {

	g := gin.New()
	g.Use(
		gin.Recovery(),
	)

	// cors
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	commit := ""
	if config.IsShowCommit {

		buildDescription := readBuildDescription()
		commit = buildDescription.Commit
	}

	g.GET("/", func(ctx *gin.Context) {
		res := map[string]string{
			"service":   config.ServiceName,
			"timestamp": time.Now().String(),
			"commit":    commit,
		}

		ctx.JSON(http.StatusOK, res)
	})

	pprof.Register(g, "dev/pprof")

	return &defaultGin{
		server: g,
	}
}

func (c *defaultGin) Engine() *gin.Engine {
	return c.server
}
