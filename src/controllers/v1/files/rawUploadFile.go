package files

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/configs/types"
	"video-conversion-service/src/services"
)

// UploadRawFile @BasePath /api/v1
// @Tags Files
// RawUpload godoc
// @Summary upload any file in raw format
// @Schemes
// @Description execution will upload any file in raw format
// @Param file formData types.FileJson true "request"
// @Param file formData file true "File"
// @Accept multipart/form-data; boundary=normal
// @Produce json
// @Success 200  {object} types.ResponseObject
// @Error 400  {object} types.ErrorObject
// @Router /files/raw [post]
func UploadRawFile(gtx *gin.Context) {
	var form types.FileJson
	if bindError := gtx.ShouldBind(&form); bindError != nil {
		funtions.ErrorResponse(gtx, "File upload mulfuntion", bindError.Error())
		return
	}

	file, err := gtx.FormFile("file")
	if err != nil {
		funtions.ErrorResponse(gtx, "File form not found", err.Error())
		return
	}

	fileName := gtx.PostForm("name")
	folder := gtx.PostForm("folder")

	mountPath := gtx.PostForm("mountPath")
	destinationPath := funtions.MakeDirSync(mountPath, folder)

	uploadedFileName, err1 := services.SaveFileToDir(gtx, file, fileName, destinationPath)
	if err1 != nil {
		funtions.ErrorResponse(gtx, "File upload error", err1.Error())
		return
	}

	funtions.SuccessResponse(gtx, "File uploaded successfully", 1, gin.H{
		"filename": destinationPath + "/" + uploadedFileName,
		"size":     file.Size,
	})
}
