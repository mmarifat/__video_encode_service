package services

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"video-conversion-service/src/configs/funtions"
)

func SaveFileToDir(gtx *gin.Context, file *multipart.FileHeader, fileName string, destinationPath string) (string, error) {
	extension := filepath.Ext(file.Filename)

	if len(fileName) < 1 {
		fileName = strings.TrimSuffix(file.Filename, extension)
	}

	fileName = regexp.MustCompile(`\s+`).ReplaceAllString(fileName, `-`)
	fileName += "-" + strconv.Itoa(int(funtions.MakeTimestamp()))
	fileName += extension
	destinationPath += "/" + fileName

	if err := gtx.SaveUploadedFile(file, destinationPath); err != nil {
		return "", err
	}

	return fileName, nil
}
