package routes

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/controllers/v1/default"
)

func DefaultRoutes(router *gin.RouterGroup) {
	defaultRoute := router.Group("/default")
	{
		defaultRoute.POST("/file", _default.UploadFile)
	}
}
