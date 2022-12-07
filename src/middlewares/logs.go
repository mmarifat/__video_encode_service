package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"video-conversion-service/src/configs/funtions"
)

func LogsMiddleware() gin.HandlerFunc {
	mode := funtions.DotEnvVariable("GIN_MODE")

	if mode == "debug" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
		// color aren't needed in release mode
		gin.DisableConsoleColor()
		// in release mode, use the file to store the success and error logs
		// loggerSuccessFile, _ := os.Create("logs/success.log")
		// loggerErrorFile, _ := os.Create("logs/error.log")
		// gin.DefaultWriter = loggerSuccessFile
		// gin.DefaultErrorWriter = loggerErrorFile
	}
	return gin.LoggerWithConfig(
		gin.LoggerConfig{
			Formatter: func(param gin.LogFormatterParams) string {
				return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s \"%s\" %s\"\n",
					param.TimeStamp.Format(time.RFC1123),
					param.ClientIP,
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.Request.UserAgent(),
					param.ErrorMessage,
				)
			},
			SkipPaths: []string{
				"/swagger/index.html",
				"/swagger/swagger-ui.css.map",
				"/swagger/swagger-ui.css",
				"/swagger/swagger-ui-bundle.js.map",
				"/swagger/swagger-ui-bundle.js",
				"/swagger/swagger-ui-standalone-preset.js.map",
				"/swagger/swagger-ui-standalone-preset.js",
				"/swagger/doc.json",
				"/swagger/favicon-32x32.png",
				"/swagger/favicon-16x16.png",
				"/swagger/favicon-64x64.png",
				"/swagger/favicon-128x128.png",
				"/swagger/favicon-256x256.png",
			},
		},
	)
}
