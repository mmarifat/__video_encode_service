package main

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"video-conversion-service/docs"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/middlewares"
	"video-conversion-service/src/routes"
)

func main() {
	port := funtions.DotEnvVariable("PORT")

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
		routes.BaseRoutes(v1Router)
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

	log.Println("APi is running at http://localhost:" + port + "\n")
	log.Println("APi Docs is running at http://localhost:" + port + "/swagger/index.html\n")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
