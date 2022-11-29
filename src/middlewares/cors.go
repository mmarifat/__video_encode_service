package middlewares

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(gtx *gin.Context) {
		gtx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		gtx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		gtx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		gtx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS, GET, PUT")

		if gtx.Request.Method == "OPTIONS" {
			gtx.AbortWithStatus(204)
			return
		}

		gtx.Next()
	}
}
