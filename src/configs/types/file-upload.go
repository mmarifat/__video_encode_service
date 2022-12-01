package types

import "mime/multipart"

type FileJson struct {
	MountPath *string               `form:"mountPath" binding:"required" default:"/mnt/streamstudio"`
	Name      *string               `form:"name"`
	Folder    *string               `form:"folder" binding:"required" default:"videoclips"`
	File      *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}

type FileCompressJson struct {
	MountPath     *string               `form:"mountPath" binding:"required" default:"/mnt/streamstudio"`
	EncodeWaiting *bool                 `form:"encodeWaiting" binding:"required" default:"false"`
	Name          *string               `form:"name"`
	Folder        *string               `form:"folder" binding:"required" default:"videoclips"`
	FfmpegStr     *string               `form:"ffmpegStr" default:"-filter:v fps=25 -vf scale=1280:720 -b:v 880k -b:a 128k -c:v h264 -c:a aac -ac 2 -ar 44100"`
	OutputFormat  *string               `form:"outputFormat" default:"mp4"`
	File          *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`
}

type FileDeleteJson struct {
	MountPathWithName string `json:"mountPathWithName" binding:"required"`
}
