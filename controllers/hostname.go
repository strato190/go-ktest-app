package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// HostnameController struct
type HostnameController struct{}

//Retrieve info about hostname
func (u HostnameController) Retrieve(c *gin.Context) {
	getHostname := getHostname()
	c.JSON(http.StatusOK, getHostname)
	return
}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}
