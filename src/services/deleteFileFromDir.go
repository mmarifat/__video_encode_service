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
