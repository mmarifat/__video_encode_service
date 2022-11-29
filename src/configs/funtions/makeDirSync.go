package funtions

import (
	"os"
	"strings"
)

func MakeDirSync(mountPath string, dirType string) string {
	if strings.HasSuffix(mountPath, "/") == false {
		mountPath += "/"
	}
	if len(dirType) > 0 {
		mountPath += dirType
	} else {
		mountPath += "files"
	}

	if _, err := os.Stat(mountPath); os.IsNotExist(err) {
		os.Mkdir(mountPath, os.ModePerm)
	}
	return mountPath
}
