package files

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"video-conversion-service/src/configs/funtions"
	"video-conversion-service/src/configs/types"
	"video-conversion-service/src/services"
)

// UploadCompressFile @BasePath /api/v1
// @Tags Files
// CompressUpload godoc
// @Summary upload any file in compress format
// @Schemes
// @Description execution will upload any file in compress format
// @Param file formData types.FileCompressJson true "request"
// @Param file formData file true "File"
// @Accept multipart/form-data; boundary=normal
// @Produce json
// @Success 200 {object} types.ResponseObject
// @Failure 400 {object} types.ErrorObject
// @Router /files/compress [post]
func UploadCompressFile(gtx *gin.Context) {
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

	mountPath := gtx.PostForm("mountPath")
	destinationPath := funtions.MakeDirSync(mountPath)

	fileName := gtx.PostForm("name")
	uploadedFileName, err1 := services.SaveFileToDir(gtx, file, fileName, destinationPath)
	if err1 != nil {
		funtions.ErrorResponse(gtx, "File upload error", err1.Error())
		return
	}

	ffmpegStr := gtx.PostForm("ffmpegStr")
	readAtNativeFrame := gtx.PostForm("readAtNativeFrame")
	outputFormat := gtx.PostForm("outputFormat")
	fileNameWithFfmpeg := services.GenerateFfmpegFileName(uploadedFileName, outputFormat)

	apiResponseMessage := "File uploaded and encoded successfully"
	if gtx.PostForm("encodeWaiting") == "true" {
		err2 := services.SaveWithFfmpegTool(destinationPath, uploadedFileName, fileNameWithFfmpeg, ffmpegStr, readAtNativeFrame)
		if err2 != nil {
			log.Println("File encoding of " + fileNameWithFfmpeg + "error " + err2.Error())
		}
	} else {
		apiResponseMessage = "File uploaded and put in encoding queue successfully"
		go func() {
			err2 := services.SaveWithFfmpegTool(destinationPath, uploadedFileName, fileNameWithFfmpeg, ffmpegStr, readAtNativeFrame)
			if err2 != nil {
				log.Println("File encoding of " + fileNameWithFfmpeg + "error " + err2.Error())
			}
		}()
	}
	funtions.SuccessResponse(gtx, apiResponseMessage, 1, gin.H{
		"filename":     strings.ReplaceAll(fileNameWithFfmpeg, "-----enc.", "."),
		"orifinalSize": file.Size,
	})
}
