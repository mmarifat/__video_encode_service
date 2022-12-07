package services

import (
	"os"
	"path/filepath"
	"strings"
)

func RenameFfmpegFileToOriginal(inputFileWithDest string, ffmpegFileWithDest string) (string, error) {
	generatedFile := strings.TrimSuffix(inputFileWithDest, filepath.Ext(inputFileWithDest))
	generatedFile += filepath.Ext(ffmpegFileWithDest)

	_, err := DeleteFileFromDir(inputFileWithDest)
	if err != nil {
		return "", err
	}
	if osErr := os.Rename(ffmpegFileWithDest, generatedFile); osErr != nil {
		return "", osErr
	}
	return generatedFile, nil
}
