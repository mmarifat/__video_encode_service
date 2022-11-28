package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
	"strings"
	"video-conversion-service/docs"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/controllers/v1"
	"video-conversion-service/src/middlewares"
	"video-conversion-service/src/routes"
)

func main() {
	port := funtions.DotEnvVariable("PORT")
	// kill previously invoked port (need for development only for HMR)
	funtions.KillPort(port)

	router := gin.New()
	// router.MaxMultipartMemory = 1024 << 20 // 1024 MB (Max memory to be used when uploading file)
	basePath := "/api/v1"

	// middlwaers
	router.Use(
		middlewares.LogsMiddleware(),
		gin.Recovery(),
		middlewares.CORSMiddleware(),
	)

	// group routes
	v1Router := router.Group(basePath)
	{
		v1Router.GET("/status", v1.ApiStatus)
		routes.RawRoutes(v1Router)
		routes.CompressRoutes(v1Router)
	}

	// By raw, http.ListenAndServe (which gin.Run wraps) will serve an unbounded number of requests.
	// Limiting the number of simultaneous connections can sometimes greatly speed things up under load.
	if funtions.DotEnvVariable("LIMIT_NO_OF_CONNECTION") == "true" {
		maxNoOfConnection, err := strconv.Atoi(funtions.DotEnvVariable("LIMIT_MAX_NO_OF_CONNECTION"))
		if err == nil {
			v1Router.Use(middlewares.MaxAllowedConnection(maxNoOfConnection))
		}
	}
	// rate limiter
	if funtions.DotEnvVariable("LIMIT_RATE") == "true" {
		middlewares.RateLimiter(v1Router)
	}

	// Swagger API docs
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Title = "Video Conversion Service"
	docs.SwaggerInfo.Description = "This microservice deals with the upload and optinally compress files with the help of FFMPEG library as quality passed by user"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	fmt.Printf("APi is running at http://localhost:" + port + "\n")
	fmt.Printf("APi Docs is running at http://localhost:" + port + "/swagger/index.html\n")

	err := router.Run("localhost:" + port)

	if err != nil {
		if strings.ContainsAny(err.Error(), "bind: address already in use") == false {
			fmt.Printf("Failed to start the server as %s", err.Error())
		}
	}
}
