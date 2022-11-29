package compress

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/configs/types"
	"video-conversion-service/src/services"
)

// UploadFile @BasePath /api/v1
// @Tags COMPRESS
// CompressUpload godoc
// @Summary upload any file in compress format
// @Schemes
// @Description execution will upload any file in compress format
// @Param file formData types.FileCompressJson true "request"
// @Param file formData file true "File"
// @Accept multipart/form-data; boundary=normal
// @Produce json
// @Success 200  {object} types.ResponseObject
// @Error 400  {object} types.ErrorObject
// @Router /compress/file [post]
func UploadFile(gtx *gin.Context) {
	var form types.FileCompressJson
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
	dirType := gtx.PostForm("type")

	mountPath := gtx.PostForm("mountPath")
	destinationPath := funtions.MakeDirSync(mountPath, dirType)

	uploadedFileName, err1 := services.SaveFileToDir(gtx, file, fileName, destinationPath)
	if err1 != nil {
		funtions.ErrorResponse(gtx, "File upload error", err1.Error())
		return
	}

	ffmpegStr := gtx.PostForm("ffmpegStr")
	outputFormat := gtx.PostForm("outputFormat")
	encodedFileName, err2 := services.SaveWithFfmpegTool(uploadedFileName, destinationPath, ffmpegStr, outputFormat)
	if err2 != nil {
		funtions.ErrorResponse(gtx, "File encoding error", err2.Error())
		return
	}

	funtions.SuccessResponse(gtx, "File uploaded and encoded successfully", 1, gin.H{
		"fileName":     destinationPath + "/" + encodedFileName,
		"orifinalSize": file.Size,
	})
}
