package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Some      string   `json:"some"`
	Variables string   `json:"variables"`
	At        []string `json:"at"`
	Config    []string `json:"config"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/config", func(c *gin.Context) {
		configPath := getEnv("APP_COINFIG", "config.yml")
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		var config Config
		err = yaml.Unmarshal([]byte(data), &config)

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		result := Config{
			Some:      config.Some,
			Variables: config.Variables,
			At:        config.At,
			Config:    config.Config,
		}

		c.JSON(http.StatusOK, result)
	})

	r.GET("/env", func(c *gin.Context) {
		c.JSON(http.StatusOK, getEnvSlice())
	})

	r.GET("/headers", func(c *gin.Context) {
		requestHeaders := c.Request.Header
		c.JSON(http.StatusOK, requestHeaders)
	})

	r.GET("/", func(c *gin.Context) {
		responseMessage := fmt.Sprintf("app is running on host: %s", getHostname())
		c.String(http.StatusOK, responseMessage)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

func getEnvSlice() []string {
	var envSlice []string
	for _, pair := range os.Environ() {
		envSlice = append(envSlice, pair)
	}
	return envSlice
}
