package main

import (
	"fmt"
	cpuCtrl "github.com/appleboy/gin-status-api"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"wtv-video-conversion-service/configs"
	"wtv-video-conversion-service/middlewares"
	"wtv-video-conversion-service/routes"
)

func main() {
	port := configs.DotEnvVariable("PORT")
	mode := configs.DotEnvVariable("GIN_MODE")

	// kill previously invoked port (need for development only for HMR)
	configs.KillPort(port)

	if mode == "debug" {
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
		// color aren't needed in release mode
		gin.DisableConsoleColor()
		// in release mode, use the file to store the success and error logs
		loggerSuccessFile, _ := os.Create("logs/success.log")
		loggerErrorFile, _ := os.Create("logs/error.log")
		gin.DefaultWriter = loggerSuccessFile
		gin.DefaultErrorWriter = loggerErrorFile
	}

	router := gin.New()

	// By default, http.ListenAndServe (which gin.Run wraps) will serve an unbounded number of requests.
	// Limiting the number of simultaneous connections can sometimes greatly speed things up under load.
	if configs.DotEnvVariable("LIMIT_NO_OF_CONNECTION") == "true" {
		maxNoOfConnection, err := strconv.Atoi(configs.DotEnvVariable("LIMIT_MAX_NO_OF_CONNECTION"))
		if err == nil {
			router.Use(middlewares.MaxAllowedConnection(maxNoOfConnection))
		}
	}
	// rate limiter
	if configs.DotEnvVariable("LIMIT_RATE") == "true" {
		middlewares.RateLimiter(router)
	}

	// upload max limit from env or maximum 10
	maxMegaByte, __err := strconv.ParseInt(configs.DotEnvVariable("MAX_UPLOAD_SIZE_MEGABYTE"), 10, 64)
	if __err != nil {
		router.MaxMultipartMemory = maxMegaByte << 20
	} else {
		router.MaxMultipartMemory = 10 << 20 // 10 MB
	}

	// middlwaers
	router.Use(middlewares.LogsMiddleware())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	// group routes
	v1Router := router.Group("/v1")
	{
		v1Router.GET("/status", cpuCtrl.GinHandler)
		routes.DefaultRoutes(v1Router)
		routes.CompressRoutes(v1Router)
	}

	err := router.Run("localhost:" + port)
	if err != nil {
		fmt.Printf("Failed to start the server")
	}
}
