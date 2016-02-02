package main

import (
	"io"
	"os"

	"github.com/freeusd/solebtc/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func init() {
	initConfig()
}

func main() {
	router := gin.New()

	var (
		logWriter   io.Writer = os.Stdout
		panicWriter io.Writer = os.Stderr
	)

	router.Use(gin.LoggerWithWriter(logWriter))
	router.Use(gin.RecoveryWithWriter(panicWriter))
	router.Use(gin.ErrorLoggerT(gin.ErrorTypeAny))

	router.Run(config.HTTP.Port)
}
