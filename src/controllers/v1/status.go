package v1

import (
	cpuCtrl "github.com/appleboy/gin-status-api"
	"github.com/gin-gonic/gin"
)

// ApiStatus @BasePath /api/v1
// @Tags status
// StatusExample godoc
// @Summary returns gin and cpu status
// @Schemes
// @Description execution will return gin and cpu status
// @Accept json
// @Produce json
// @Success 200  {object} types.ResponseObject
// @Error 400  {object} types.ErrorObject
// @Router /status [get]
func ApiStatus(c *gin.Context) {
	cpuCtrl.GinHandler(c)
}
