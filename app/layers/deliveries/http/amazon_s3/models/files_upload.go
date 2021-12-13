package models

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"upload-file/app/entities"
	"upload-file/app/errors"
)

type UploadFilesRequestFile struct {
	FileHeader *multipart.FileHeader `form:"file" binding:"required"`
	Files      multipart.File
}

type UploadFileResponse struct {
	Url string `json:"url"`
}

func (model *UploadFileResponse) Parse(data *entities.S3Response) *UploadFileResponse {
	model.Url = data.Location
	return model
}

func (model *UploadFilesRequestFile) Parse(c *gin.Context) (*UploadFilesRequestFile, error) {
	err := c.ShouldBind(model)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	model.Files, _, _ = c.Request.FormFile("file")
	return model, nil
}

func (model *UploadFilesRequestFile) ToEntity() entities.FilesUpload {
	filesEntity := entities.FilesUpload{
		Files:       model.Files,
		FilesHeader: *model.FileHeader,
	}
	return filesEntity
}
