package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"video-conversion-service/src/configs/funtions"
)

func MaxUploadBodySizeMiddleware() gin.HandlerFunc {
	return func(gtx *gin.Context) {
		var w http.ResponseWriter = gtx.Writer
		// upload max limit from env or maximum 10
		maxMegaByte, __err := strconv.ParseInt(funtions.DotEnvVariable("MAX_UPLOAD_SIZE_MEGABYTE"), 10, 64)
		if __err == nil {
			gtx.Request.Body = http.MaxBytesReader(w, gtx.Request.Body, 1024*1024*maxMegaByte)
		} else {
			// Default 100 MegaByte
			gtx.Request.Body = http.MaxBytesReader(w, gtx.Request.Body, 1024*1024*100)
		}
		gtx.Next()
	}
}
