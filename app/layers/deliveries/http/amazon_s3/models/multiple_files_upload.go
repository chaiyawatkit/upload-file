package models

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"mime/multipart"
	"path/filepath"
	"sort"
	"strings"
	"upload-file/app/entities"
	"upload-file/app/errors"
)

type UploadMultipleFilesRequestFile struct {
	FileHeader []multipart.FileHeader `form:"file" binding:"required"`
}

type ResponseJSON struct {
	Location  string  `json:"Location"`
}

type MultipleFileResponseJSON []ResponseJSON

func (model *UploadMultipleFilesRequestFile) Parse(c *gin.Context) (*UploadMultipleFilesRequestFile, error) {
	err := c.ShouldBind(model)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return model, nil
}

func (model *UploadMultipleFilesRequestFile) IsValid(c *gin.Context) (*UploadMultipleFilesRequestFile, error) {

	mB := 3000000
	fileCheck := []string{".jpg", ".pdf", ".png", ".gif", ".jpeg"}
	sort.Strings(fileCheck)
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		fileType := strings.TrimSuffix(filepath.Ext(file.Filename), file.Filename)

		x := sort.SearchStrings(fileCheck, fileType)
		if x < len(fileCheck) && fileCheck[x] == fileType {
		} else {
			return nil, errors.RecordNotFoundError{"fileTypeError"}
		}
		sizeFileCal := float64(file.Size) / float64(mB)
		if sizeFileCal > 1 {
			return nil, errors.RecordNotFoundError{"fileSizeError"}
		}
	}
	return model, nil
}

func (model *UploadMultipleFilesRequestFile) ToEntity() entities.MutipleUpload {
	filesEntity := entities.MutipleUpload{
		FilesHeader: model.FileHeader,
	}
	return filesEntity
}

func (model *MultipleFileResponseJSON) ResponseJSON(data []s3manager.UploadOutput) *MultipleFileResponseJSON {
	_ = copier.Copy(model, data)
	return model
}
