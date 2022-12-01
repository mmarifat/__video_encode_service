package v1

import (
	golangStatsApi "github.com/fukata/golang-stats-api-handler"
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
)

// ApiStatus @BasePath /api/v1
// @Tags Status
// ApiStatus godoc
// @Summary returns gin and cpu status
// @Schemes
// @Description execution will return gin and cpu status
// @Accept json
// @Produce json
// @Success 200 {object} types.ResponseObject
// @Failure 400 {object} types.ErrorObject
// @Router /status [get]
func ApiStatus(gtx *gin.Context) {
	funtions.SuccessResponse(gtx, "Service Status", 1, golangStatsApi.GetStats())
}
