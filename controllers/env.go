package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//EnvsController struct
type EnvsController struct{}

//Retrieve info about environment variables
func (u EnvsController) Retrieve(c *gin.Context) {
	envSlice := getEnvSlice()
	c.JSON(http.StatusOK, envSlice)
	return
}

func getEnvSlice() []string {
	var envSlice []string
	for _, pair := range os.Environ() {
		envSlice = append(envSlice, pair)
	}
	return envSlice
}
