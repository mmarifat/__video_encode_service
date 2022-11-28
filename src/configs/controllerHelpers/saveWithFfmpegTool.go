package controllerHelpers

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"video-conversion-service/src/configs/funtions"
)

func SaveWithFfmpegTool(fileName string, dirType string, ffmpegString string, outputFormat string) (string, error) {
	destination := funtions.MakeDirSync(dirType)

	inputFile := destination + "/" + fileName

	extension := filepath.Ext(fileName)
	destinationWithFfmpeg := strings.TrimSuffix(fileName, extension)
	if len(outputFormat) > 0 {
		destinationWithFfmpeg += "--enc." + outputFormat
	} else {
		destinationWithFfmpeg += "--enc" + extension
	}
	destinationWithFfmpegFile := destination + "/" + destinationWithFfmpeg

	ffmpegStringArgs := `ffmpeg -i ` + inputFile + ` ` + ffmpegString + ` ` + destinationWithFfmpegFile
	ffmpegStringArgs = regexp.MustCompile(`\s+`).ReplaceAllString(ffmpegStringArgs, ` `)
	ffmpegStringArgs = strings.ReplaceAll(ffmpegStringArgs, "\"", "'")
	args := strings.Split(ffmpegStringArgs, ` `)
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	// remove previous original file
	if osErr := os.Remove(inputFile); osErr != nil {
		return "", osErr
	}

	return destinationWithFfmpeg, nil
}
