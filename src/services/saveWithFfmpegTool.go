package services

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

	if len(ffmpegString) < 1 {
		//Width: 1280
		//Height: 720
		//Video Bitrate: 880Kbps
		//Framerate: 25FPS
		//Container: MP4
		//Audio Bitrate: 128Kbps
		//Audio: AAC
		//Channels: 2
		//Audio Sample Rate: 44.100kHz
		ffmpegString = "-filter:v fps=25 -vf scale=1280:720 -b:v 880k -b:a 128k -c:v h264 -c:a aac -ac 2 -ar 44100"
	}
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
