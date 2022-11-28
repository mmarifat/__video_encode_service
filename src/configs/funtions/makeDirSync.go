package funtions

import "os"

func MakeDirSync(dirType string) string {
	destination := "uploaded/"
	if len(dirType) > 0 {
		destination += dirType
	} else {
		destination += "files"
	}

	if _, err := os.Stat(destination); os.IsNotExist(err) {
		os.Mkdir(destination, os.ModePerm)
	}
	return destination
}
