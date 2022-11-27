package controllers

import (
	cpuCtrl "github.com/appleboy/gin-status-api"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// StatusExample godoc
// @Summary returns gin and cpu status
// @Schemes
// @Description execution will return gin and cpu status
// @Tags status
// @Accept json
// @Produce json
// @Success 200
// @Error 400
// @Router /status [get]
func ApiStatus(c *gin.Context) {
	cpuCtrl.GinHandler(c)
}
