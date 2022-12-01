package funtions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"video-conversion-service/src/configs/types"
)

func SuccessResponse(gtx *gin.Context, msg string, count int, data any) {
	resp := &types.ResponseObject{Nonce: MakeTimestamp(), Statuscode: http.StatusOK, Message: msg, Payload: gin.H{
		"count": count,
		"data":  data,
	}}
	gtx.Header("Content-Type", "application/json")
	gtx.JSON(resp.Statuscode, resp)
}

func ErrorResponse(gtx *gin.Context, msg string, error any) {
	resp := &types.ErrorObject{Nonce: MakeTimestamp(), Statuscode: http.StatusBadRequest, Message: msg, Error: error}
	gtx.Header("Content-Type", "application/json")
	gtx.JSON(resp.Statuscode, resp)
}
