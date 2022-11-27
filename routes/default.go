package routes

import (
	"github.com/gin-gonic/gin"
	"wtv-video-conversion-service/configs"
)

func DefaultRoutes(router *gin.RouterGroup) {
	defaultRoute := router.Group("/default")
	{
		defaultRoute.GET("/picture", func(c *gin.Context) {
			configs.SuccessResponse(c, "pong", gin.H{
				"count": 1,
				"data":  "ok default",
			})
		})
	}
}
