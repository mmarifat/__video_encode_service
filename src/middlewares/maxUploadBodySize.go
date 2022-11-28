package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"video-conversion-service/src/configs/funtions"
)

func MaxUploadBodySizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("im here %s", c)
		var w http.ResponseWriter = c.Writer
		// upload max limit from env or maximum 10
		maxMegaByte, __err := strconv.ParseInt(funtions.DotEnvVariable("MAX_UPLOAD_SIZE_MEGABYTE"), 10, 64)
		if __err == nil {
			c.Request.Body = http.MaxBytesReader(w, c.Request.Body, 1024*1024*maxMegaByte)
		} else {
			// Default 100 MegaByte
			c.Request.Body = http.MaxBytesReader(w, c.Request.Body, 1024*1024*100)
		}
		c.Next()
	}
}
