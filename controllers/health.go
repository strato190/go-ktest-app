package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/strato190/go-ktest-app/util/logging"
)

//HealthController struct
type HealthController struct{}

//Retrieve info about app health. to be improved
func (u HealthController) Retrieve(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	//logging.Info("health endpoint invoked")
	return
}
