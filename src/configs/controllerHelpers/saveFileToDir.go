package controllerHelpers

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"video-conversion-service/src/configs/funtions"
)

func SaveFileToDir(c *gin.Context, file *multipart.FileHeader, fileName string, dirType string) string {
	destination := "uploaded/"
	if len(dirType) > 0 {
		destination += dirType
	} else {
		destination += "files"
	}

	if _, err := os.Stat(destination); os.IsNotExist(err) {
		os.Mkdir(destination, os.ModePerm)
	}

	extension := filepath.Ext(file.Filename)

	if len(fileName) < 1 {
		fileName = strings.TrimSuffix(file.Filename, extension)
	}

	fileName += "-" + strconv.Itoa(int(funtions.MakeTimestamp()))
	fileName += extension
	destination += "/" + fileName

	if err := c.SaveUploadedFile(file, destination); err != nil {
		funtions.ErrorResponse(c, "File upload error", err.Error())
		return ""
	}

	return fileName
}
