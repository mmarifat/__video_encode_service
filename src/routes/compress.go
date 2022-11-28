package routes

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/controllers/v1/compress"
	"video-conversion-service/src/middlewares"
)

func CompressRoutes(router *gin.RouterGroup) {
	compressRoute := router.Group("/compress")
	{
		compressRoute.POST("/file", middlewares.MaxUploadBodySizeMiddleware(), compress.UploadFile)
	}
}
