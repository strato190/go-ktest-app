package server

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"

	ginprometheus "github.com/banzaicloud/go-gin-prometheus"
	"github.com/strato190/go-ktest-app/controllers"
	"github.com/strato190/go-ktest-app/middleware"
	//_ "github.com/strato190/go-ktest-app/docs"
	//"github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

func stop(c *gin.Context) {
	syscall.Kill(os.Getpid(), syscall.SIGTERM)
}

func setupRouter() *gin.Engine {
	r := gin.New()
	p := ginprometheus.NewPrometheus("gin", []string{})
	p.SetListenAddress(":9781")
	p.Use(r, "/metrics")

	// Default logger from gin
	//r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(middleware.LogMiddleware())
	r.Use(middleware.RequestID(middleware.RequestIDOptions{AllowSetting: false}))

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
		sc := new(controllers.SleepController)
		v1.GET("/sleep", sc.Retrieve)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "app is up and running ")
	})

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/stop", stop)

	return r
}
