package configs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseObject struct {
	Nonce      int    `json:"nonce"`
	Statuscode int    `json:"status"`
	Message    string `json:"msg"`
	Payload    any    `json:"payload"`
}

func SuccessResponse(c *gin.Context, msg string, data any) {
	resp := &ResponseObject{Statuscode: http.StatusOK, Message: msg, Payload: data}
	c.Header("Content-Type", "application/json")
	c.JSON(resp.Statuscode, resp)
}

func ErrorResponse(c *gin.Context, msg string, data any) {
	resp := &ResponseObject{Statuscode: http.StatusBadRequest, Message: msg, Payload: data}
	c.Header("Content-Type", "application/json")
	c.JSON(resp.Statuscode, resp)
}
