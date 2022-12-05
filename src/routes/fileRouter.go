package routes

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/controllers/v1/files"
	"video-conversion-service/src/middlewares"
)

func FileRoutes(router *gin.RouterGroup) {
	fileRouter := router.Group("/files")
	{
		fileRouter.POST("/raw", middlewares.MaxUploadBodySizeMiddleware(), files.UploadRawFile)
		fileRouter.POST("/compress", middlewares.MaxUploadBodySizeMiddleware(), files.UploadCompressFile)
		fileRouter.DELETE("/remove", files.DeleteFile)
	}
}
