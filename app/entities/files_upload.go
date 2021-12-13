package entities

import "mime/multipart"

type FilesUpload struct {
	Files       multipart.File
	FileName    *string
	FilesHeader multipart.FileHeader
}

type S3Response struct {
	Location  string
	VersionID string
	UploadID  string
	ETag      string
}
