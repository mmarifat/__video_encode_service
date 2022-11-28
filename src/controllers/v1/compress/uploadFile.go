package compress

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/controllerHelpers"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/configs/types"
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

	file, err1 := c.FormFile("file")
	if err1 != nil {
		funtions.ErrorResponse(c, "File form not found", err1.Error())
		return
	}

	fileName := c.PostForm("name")
	dirType := c.PostForm("type")
	uploadedFileName, err2 := controllerHelpers.SaveFileToDir(c, file, fileName, dirType)
	if err2 != nil {
		funtions.ErrorResponse(c, "File upload error", err2.Error())
		return
	}

	ffmpegStr := c.PostForm("ffmpegStr")
	encodeInfo, err3 := controllerHelpers.SaveWithFfmpegTool(uploadedFileName, dirType, ffmpegStr)
	if err3 != nil {
		funtions.ErrorResponse(c, "File encoding error", err3.Error())
		return
	}

	funtions.SuccessResponse(c, "File uploaded and encoded successfully", gin.H{
		"count": 1,
		"data": gin.H{
			"output":       encodeInfo,
			"orifinalSize": file.Size,
		},
	})
}
