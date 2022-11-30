package v1

import (
	cpuCtrl "github.com/appleboy/gin-status-api"
	"github.com/gin-gonic/gin"
)

// ApiStatus @BasePath /api/v1
// @Tags Status
// getStatus godoc
// @Summary returns gin and cpu status
// @Schemes
// @Description execution will return gin and cpu status
// @Accept json
// @Produce json
// @Success 200  {object} object
// @Error 400  {object} object
// @Router /status [get]
func ApiStatus(gtx *gin.Context) {
	cpuCtrl.GinHandler(gtx)
}
