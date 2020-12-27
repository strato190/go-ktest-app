package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//SleepController struct
type SleepController struct{}

//Retrieve info about environment variables
func (u SleepController) Retrieve(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	time.Sleep(120 * time.Second)
	return
}
