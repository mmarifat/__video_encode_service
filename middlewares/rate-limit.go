package middlewares

import (
	"fmt"
	"github.com/axiaoxin-com/goutils"
	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
	"video-conversion-service/configs"
)

func RateLimiter(router *gin.Engine) {
	// Put a token into the token bucket every 1s
	// Maximum 1 request allowed per second
	rdb, err := goutils.NewRedisClient(&redis.Options{})
	if err != nil {
		panic(err)
	}
	router.Use(ratelimiter.GinRedisRatelimiter(rdb, ratelimiter.GinRatelimiterConfig{
		// config: rate limiter key using client IP Address
		LimitKey: func(c *gin.Context) string {
			return c.ClientIP()
		},
		// config: how to respond when limiting
		LimitedHandler: func(c *gin.Context) {
			configs.ErrorResponse(c, "Too many requests", nil)
			c.Abort()
			return
		},
		// config: return ratelimiter token fill interval and bucket size (every 1 second)
		TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
			intervalSecond, err1 := strconv.Atoi(configs.DotEnvVariable("LIMIT_RATE_INTERVAL_SECOND"))
			bucketSize, err2 := strconv.Atoi(configs.DotEnvVariable("LIMIT_RATE_BUCKET_SIZE"))
			if err1 == nil && err2 == nil {
				fmt.Printf("ok %s", time.Second*time.Duration(intervalSecond))
				return time.Second * time.Duration(intervalSecond), bucketSize
			}
			// else 1 second with 1 bucket size
			fmt.Printf("test %s", time.Second*1)
			return time.Second * 1, 1
		},
	}))
}
