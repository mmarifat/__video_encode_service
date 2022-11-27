package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"strconv"
	"video-conversion-service/docs"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/controllers/v1"
	middlewares2 "video-conversion-service/src/middlewares"
	routes2 "video-conversion-service/src/routes"
)

func main() {
	port := funtions.DotEnvVariable("PORT")
	mode := funtions.DotEnvVariable("GIN_MODE")

	// kill previously invoked port (need for development only for HMR)
	funtions.KillPort(port)

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
	// router.MaxMultipartMemory = 10 << 20 // 10 MB (Max memory to be used when uploading file)
	basePath := "/api/v1"

	// middlwaers
	router.Use(middlewares2.LogsMiddleware())
	router.Use(gin.Recovery())
	router.Use(middlewares2.CORSMiddleware())
	router.Use(middlewares2.MaxUploadBodySizeMiddleware())

	// group routes
	v1Router := router.Group(basePath)
	{
		v1Router.GET("/status", v1.ApiStatus)
		routes2.DefaultRoutes(v1Router)
		routes2.CompressRoutes(v1Router)
	}

	// By default, http.ListenAndServe (which gin.Run wraps) will serve an unbounded number of requests.
	// Limiting the number of simultaneous connections can sometimes greatly speed things up under load.
	if funtions.DotEnvVariable("LIMIT_NO_OF_CONNECTION") == "true" {
		maxNoOfConnection, err := strconv.Atoi(funtions.DotEnvVariable("LIMIT_MAX_NO_OF_CONNECTION"))
		if err == nil {
			v1Router.Use(middlewares2.MaxAllowedConnection(maxNoOfConnection))
		}
	}
	// rate limiter
	if funtions.DotEnvVariable("LIMIT_RATE") == "true" {
		middlewares2.RateLimiter(v1Router)
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
