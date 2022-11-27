package types

import "mime/multipart"

type FileJson struct {
	Name string                `form:"name"`
	Type string                `form:"type" binding:"required" default:"picture"`
	File *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}
