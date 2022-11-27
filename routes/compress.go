package routes

import (
	"github.com/gin-gonic/gin"
	"wtv-video-conversion-service/configs"
)

func CompressRoutes(router *gin.RouterGroup) {
	compressRoute := router.Group("/compress")
	{
		compressRoute.GET("/picture", func(c *gin.Context) {
			configs.SuccessResponse(c, "pong", gin.H{
				"count": 1,
				"data":  "ok",
			})
		})
	}
}
