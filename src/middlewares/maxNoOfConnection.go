package middlewares

import "github.com/gin-gonic/gin"

func MaxAllowedConnection(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(gtx *gin.Context) {
		acquire()       // before request
		defer release() // after request
		gtx.Next()

	}
}
