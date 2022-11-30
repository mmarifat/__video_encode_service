package routes

import (
	"github.com/gin-gonic/gin"
	v1 "video-conversion-service/src/controllers/v1"
)

func BaseRoutes(router *gin.RouterGroup) {
	router.GET("/status", v1.ApiStatus)
	FileRoutes(router)
}
