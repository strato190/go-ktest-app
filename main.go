package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/strato190/go-ktest-app/server"
)

//Version variable. Placeholder for ldflag replacement
var Version = "0.0.0"

//GitCommit variable. Placeholder for ldflag replacement
var GitCommit = "aaaaaaa"

//BuildTime variable. Placeholder for ldflag replacement
var BuildTime = "1970-01-01 00:00:00"

func main() {
	log.Printf(fmt.Sprintf("Starting app version %s with commit sha: %s built on %s", Version, GitCommit, BuildTime))

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	server.Init()
}
