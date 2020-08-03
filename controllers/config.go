package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

//Config struct
type Config struct {
	Some      string   `json:"some"`
	Variables string   `json:"variables"`
	At        []string `json:"at"`
	Config    []string `json:"config"`
}

//ConfigController struct
type ConfigController struct{}

//Retrieve info about config file
func (u ConfigController) Retrieve(c *gin.Context) {
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
	return
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
