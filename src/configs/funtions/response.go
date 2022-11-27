package funtions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"video-conversion-service/src/configs/types"
)

func SuccessResponse(c *gin.Context, msg string, data any) {
	resp := &types.ResponseObject{Statuscode: http.StatusOK, Message: msg, Payload: data}
	c.Header("Content-Type", "application/json")
	c.JSON(resp.Statuscode, resp)
}

func ErrorResponse(c *gin.Context, msg string, error any) {
	resp := &types.ErrorObject{Statuscode: http.StatusBadRequest, Message: msg, Error: error}
	c.Header("Content-Type", "application/json")
	c.JSON(resp.Statuscode, resp)
}
