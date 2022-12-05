package files

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/services"
)

// DeleteFile @BasePath /api/v1
// @Tags Files
// deleteFile godoc
// @Summary delete a file from specific location
// @Schemes
// @Description execution will delete a file from specific location
// @Param mountPathWithName query string true "File Name with the full mounted path"
// @Accept json
// @Produce json
// @Success 200 {object} types.ResponseObject
// @Failure 400 {object} types.ErrorObject
// @Router /files/remove [delete]
func DeleteFile(gtx *gin.Context) {
	mountPathWithName, ok := gtx.GetQuery("mountPathWithName")
	if ok != true {
		funtions.ErrorResponse(gtx, "File delete payload mulfuntion! ", nil)
		return
	}
	_, osErr := services.DeleteFileFromDir(mountPathWithName)
	if osErr != nil {
		funtions.ErrorResponse(gtx, "File delete error! ", osErr.Error())
		return
	}
	funtions.SuccessResponse(gtx, "File deleted successfully", 1, gin.H{
		"isDeleted": true,
	})
}
