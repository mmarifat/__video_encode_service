package middlewares

import (
	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"video-conversion-service/src/configs/funtions"
)

func RateLimiter(router *gin.RouterGroup) {
	router.Use(ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
		// config: rate limiter key using client IP Address
		LimitKey: func(gtx *gin.Context) string {
			return gtx.ClientIP()
		},
		// config: how to respond when limiting
		LimitedHandler: func(gtx *gin.Context) {
			funtions.ErrorResponse(gtx, "Too many requests", nil)
			gtx.Abort()
			return
		},
		// config: return ratelimiter token fill interval and bucket size (every 1 second)
		TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
			intervalSecond, err1 := strconv.Atoi(funtions.DotEnvVariable("LIMIT_RATE_INTERVAL_SECOND"))
			bucketSize, err2 := strconv.Atoi(funtions.DotEnvVariable("LIMIT_RATE_BUCKET_SIZE"))
			if err1 == nil && err2 == nil {
				return time.Second * time.Duration(intervalSecond), bucketSize
			}
			// else 1 second with 1 bucket size
			return time.Second * 1, 1
		},
	}))
}
