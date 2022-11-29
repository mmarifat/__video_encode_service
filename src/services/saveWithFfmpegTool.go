package services

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func GenerateFfmpegFileName(fileName string, outputFormat string) string {
	extension := filepath.Ext(fileName)
	generatedFileName := strings.TrimSuffix(fileName, extension)
	if len(outputFormat) > 0 {
		generatedFileName += "--enc." + outputFormat
	} else {
		generatedFileName += "--enc" + extension
	}
	return generatedFileName
}

func SaveWithFfmpegTool(fileInputWithFfmpeg string, fileDestWithFfmpeg string, ffmpegString string) (string, error) {
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
	ffmpegStringArgs := `ffmpeg -i ` + fileInputWithFfmpeg + ` ` + ffmpegString + ` ` + fileDestWithFfmpeg
	ffmpegStringArgs = regexp.MustCompile(`\s+`).ReplaceAllString(ffmpegStringArgs, ` `)
	ffmpegStringArgs = strings.ReplaceAll(ffmpegStringArgs, "\"", "'")
	args := strings.Split(ffmpegStringArgs, ` `)
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	// remove previous original file
	if osErr := os.Remove(fileInputWithFfmpeg); osErr != nil {
		return "", osErr
	}

	return fileDestWithFfmpeg, nil
}
