package entities

import "mime/multipart"

type MutipleUpload struct {
	Files       multipart.File
	FilesHeader []multipart.FileHeader
}
