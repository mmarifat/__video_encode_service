package funtions

import (
	"os"
	"strings"
)

func MakeDirSync(mountPath string) string {
	if strings.HasSuffix(mountPath, "/") == true {
		mountPath = strings.TrimSuffix(mountPath, "/")
	}
	if _, err := os.Stat(mountPath); os.IsNotExist(err) {
		os.MkdirAll(mountPath, os.ModePerm)
	}
	return mountPath
}
