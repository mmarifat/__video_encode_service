package services

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func GenerateFfmpegFileName(fileName string, outputFormat string) string {
	extension := filepath.Ext(fileName)
	generatedFileName := strings.TrimSuffix(fileName, extension)
	if len(outputFormat) > 0 {
		generatedFileName += "-----enc." + outputFormat
	} else {
		generatedFileName += "-----enc" + extension
	}
	return generatedFileName
}

func SaveWithFfmpegTool(destinationPath string, uploadedFileName string, fileNameWithFfmpeg string, ffmpegString string, readAtNativeFrame string) error {
	fileInputForFfmpeg := destinationPath + "/" + uploadedFileName
	fileDestWithFfmpeg := destinationPath + "/" + fileNameWithFfmpeg

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
	ffmpegStringArgs := `ffmpeg`
	if readAtNativeFrame == "true" {
		ffmpegStringArgs += " -re "
	}
	ffmpegStringArgs += ` -i ` + fileInputForFfmpeg + ` ` + ffmpegString + ` ` + fileDestWithFfmpeg
	ffmpegStringArgs = regexp.MustCompile(`\s+`).ReplaceAllString(ffmpegStringArgs, ` `)
	ffmpegStringArgs = strings.ReplaceAll(ffmpegStringArgs, "\"", "'")
	args := strings.Split(ffmpegStringArgs, ` `)
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.CombinedOutput()

	if err != nil {
		DeleteFileFromDir(fileDestWithFfmpeg)
		return err
	}

	// remove previous original file
	_, osErr := RenameFfmpegFileToOriginal(fileInputForFfmpeg, fileDestWithFfmpeg)
	if osErr != nil {
		return osErr
	}

	return nil
}
