package files

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/configs/types"
	"video-conversion-service/src/services"
)

// DeleteFile @BasePath /api/v1
// @Tags Files
// deleteFile godoc
// @Summary delete a file from specific location
// @Schemes
// @Description execution will delete a file from specific location
// @Param payload body types.FileDeleteJson true "Mount Path With Name"
// @Accept json
// @Produce json
// @Success 200 {object} types.ResponseObject
// @Failure 400 {object} types.ErrorObject
// @Router /files/remove [patch]
func DeleteFile(gtx *gin.Context) {
	var payload types.FileDeleteJson
	if bindError := gtx.ShouldBind(&payload); bindError != nil {
		funtions.ErrorResponse(gtx, "File delete payload mulfuntion! ", bindError.Error())
		return
	}
	_, osErr := services.DeleteFileFromDir(payload.MountPathWithName)
	if osErr != nil {
		funtions.ErrorResponse(gtx, "File delete error! ", osErr.Error())
		return
	}
	funtions.SuccessResponse(gtx, "File deleted successfully", 1, gin.H{
		"isDeleted": true,
	})
}
