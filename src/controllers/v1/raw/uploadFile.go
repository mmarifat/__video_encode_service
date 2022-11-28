package raw

import (
	"github.com/gin-gonic/gin"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/configs/types"
	"video-conversion-service/src/services"
)

// UploadFile @BasePath /api/v1
// @Tags RAW
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
// @Router /raw/file [post]
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
	uploadedFileName, err := services.SaveFileToDir(c, file, fileName, dirType)
	if err != nil {
		funtions.ErrorResponse(c, "File upload error", err.Error())
	}

	funtions.SuccessResponse(c, "File uploaded successfully", 1, gin.H{
		"filename": uploadedFileName,
		"size":     file.Size,
	})
}
