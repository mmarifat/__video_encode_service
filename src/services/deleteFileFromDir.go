package services

import (
	"os"
)

func DeleteFileFromDir(mountPathWithName string) (bool, error) {
	if osErr := os.Remove(mountPathWithName); osErr != nil {
		return false, osErr
	}
	return true, nil
}

func RenameFfmpegFileToOriginal(inputFileWithDest string, ffmpegFileWithDest string) (bool, error) {
	_, err := DeleteFileFromDir(inputFileWithDest)
	if err != nil {
		return false, err
	}
	if osErr := os.Rename(ffmpegFileWithDest, inputFileWithDest); osErr != nil {
		return false, osErr
	}
	return true, nil
}
