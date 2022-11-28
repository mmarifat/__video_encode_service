package routes

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/controllers/v1/raw"
	"video-conversion-service/src/middlewares"
)

func RawRoutes(router *gin.RouterGroup) {
	defaultRoute := router.Group("/raw")
	{
		defaultRoute.POST("/file", middlewares.MaxUploadBodySizeMiddleware(), raw.UploadFile)
	}
}
