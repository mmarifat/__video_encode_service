package compress

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/configs/types"
	services2 "video-conversion-service/src/services"
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
func UploadFile(c *gin.Context) {
	var form types.FileJson
	if bindError := c.ShouldBind(&form); bindError != nil {
		funtions.ErrorResponse(c, "File upload mulfuntion", bindError.Error())
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		funtions.ErrorResponse(c, "File form not found", err.Error())
		return
	}

	fileName := c.PostForm("name")
	dirType := c.PostForm("type")
	uploadedFileName, err := services2.SaveFileToDir(c, file, fileName, dirType)
	if err != nil {
		funtions.ErrorResponse(c, "File upload error", err.Error())
		return
	}

	ffmpegStr := c.PostForm("ffmpegStr")
	outputFormat := c.PostForm("outputFormat")
	encodedFileName, err := services2.SaveWithFfmpegTool(uploadedFileName, dirType, ffmpegStr, outputFormat)
	if err != nil {
		funtions.ErrorResponse(c, "File encoding error", err.Error())
		return
	}

	funtions.SuccessResponse(c, "File uploaded and encoded successfully", 1, gin.H{
		"fileName":     encodedFileName,
		"orifinalSize": file.Size,
	})
}
