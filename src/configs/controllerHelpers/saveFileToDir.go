package controllerHelpers

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"video-conversion-service/src/configs/funtions"
)

func SaveFileToDir(c *gin.Context, file *multipart.FileHeader, fileName string, dirType string) (string, error) {
	destination := funtions.MakeDirSync(dirType)
	extension := filepath.Ext(file.Filename)

	if len(fileName) < 1 {
		fileName = strings.TrimSuffix(file.Filename, extension)
	}

	fileName = regexp.MustCompile(`\s+`).ReplaceAllString(fileName, `-`)
	fileName += "-" + strconv.Itoa(int(funtions.MakeTimestamp()))
	fileName += extension
	destination += "/" + fileName

	if err := c.SaveUploadedFile(file, destination); err != nil {
		return "", err
	}

	return fileName, nil
}
