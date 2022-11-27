package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"strconv"
	"video-conversion-service/configs"
	"video-conversion-service/controllers"
	"video-conversion-service/docs"
	"video-conversion-service/middlewares"
	"video-conversion-service/routes"
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
	basePath := "/api/v1"

	// middlwaers
	router.Use(middlewares.LogsMiddleware())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	// group routes
	v1Router := router.Group(basePath)
	{
		v1Router.GET("/status", controllers.ApiStatus)
		routes.DefaultRoutes(v1Router)
		routes.CompressRoutes(v1Router)
	}

	// upload max limit from env or maximum 10
	maxMegaByte, __err := strconv.ParseInt(configs.DotEnvVariable("MAX_UPLOAD_SIZE_MEGABYTE"), 10, 64)
	if __err != nil {
		router.MaxMultipartMemory = maxMegaByte << 20
	} else {
		router.MaxMultipartMemory = 10 << 20 // 10 MB
	}

	// By default, http.ListenAndServe (which gin.Run wraps) will serve an unbounded number of requests.
	// Limiting the number of simultaneous connections can sometimes greatly speed things up under load.
	if configs.DotEnvVariable("LIMIT_NO_OF_CONNECTION") == "true" {
		maxNoOfConnection, err := strconv.Atoi(configs.DotEnvVariable("LIMIT_MAX_NO_OF_CONNECTION"))
		if err == nil {
			v1Router.Use(middlewares.MaxAllowedConnection(maxNoOfConnection))
		}
	}
	// rate limiter
	if configs.DotEnvVariable("LIMIT_RATE") == "true" {
		middlewares.RateLimiter(v1Router)
	}

	// Swagger API docs
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Title = "Video Conversion Service"
	docs.SwaggerInfo.Description = "This microservice deals with the upload and optinally compress files with the help of FFMPEG library as quality passed by user"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	fmt.Printf("APi Docs is running at http://localhost:" + port + "/swagger/index.html\n")

	err := router.Run("localhost:" + port)
	if err != nil {
		fmt.Printf("Failed to start the server")
	}
}
