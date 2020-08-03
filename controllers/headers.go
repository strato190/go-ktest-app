package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HeadersController struct
type HeadersController struct{}

//Retrieve info about headers
func (u HeadersController) Retrieve(c *gin.Context) {
	requestHeaders := c.Request.Header
	c.JSON(http.StatusOK, requestHeaders)
	return
}
