package funtions

import (
	"os"
	"strings"
)

func MakeDirSync(mountPath string, folder string) string {
	if strings.HasSuffix(mountPath, "/") == false {
		mountPath += "/"
	}
	if len(folder) > 0 {
		mountPath += folder
	} else {
		mountPath += "files"
	}

	if _, err := os.Stat(mountPath); os.IsNotExist(err) {
		os.Mkdir(mountPath, os.ModePerm)
	}
	return mountPath
}
