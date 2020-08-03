package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/strato190/go-ktest-app/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.New()
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
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "app is up and running ")
	})
	return r
}
