package routes

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
)

func CompressRoutes(router *gin.RouterGroup) {
	compressRoute := router.Group("/compress")
	{
		compressRoute.GET("/file", func(c *gin.Context) {
			funtions.SuccessResponse(c, "pong", gin.H{
				"count": 1,
				"data":  "ok",
			})
		})
	}
}
