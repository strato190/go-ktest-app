package server

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"

	ginprometheus "github.com/banzaicloud/go-gin-prometheus"
	"github.com/strato190/go-ktest-app/controllers"
)

func stop(c *gin.Context) {
	syscall.Kill(os.Getpid(), syscall.SIGTERM)
}

func setupRouter() *gin.Engine {
	r := gin.New()
	p := ginprometheus.NewPrometheus("gin", []string{})
	p.SetListenAddress(":9781")
	p.Use(r, "/metrics")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")
	{
		cc := new(controllers.ConfigController)
		v1.GET("/config", cc.Retrieve)
		ec := new(controllers.EnvsController)
		v1.GET("/env", ec.Retrieve)
		hc := new(controllers.HeadersController)
		v1.GET("/headers", hc.Retrieve)
		nc := new(controllers.HostnameController)
		v1.GET("/hostname", nc.Retrieve)
		lc := new(controllers.HealthController)
		v1.GET("/health", lc.Retrieve)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "app is up and running ")
	})

	r.GET("/stop", stop)

	return r
}
