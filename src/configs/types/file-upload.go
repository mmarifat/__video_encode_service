package types

import "mime/multipart"

type FileJson struct {
	Name string                `form:"name"`
	Type string                `form:"type" binding:"required" default:"files"`
	File *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}

type FileCompressJson struct {
	Name         string                `form:"name"`
	Type         string                `form:"type" binding:"required" default:"files-encoded"`
	FfmpegStr    string                `form:"ffmpegStr" binding:"required"`
	OutputFormat string                `form:"outputFormat"`
	File         *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}
