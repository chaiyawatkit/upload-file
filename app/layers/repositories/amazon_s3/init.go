package amazon_s3

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"upload-file/app/entities"
)

type repo struct {
}

func Init() Repo {
	return &repo{}
}

type Repo interface {
	FilesUpload(upload entities.FilesUpload) (*entities.S3Response, error)
	MultipleFilesUpload(upload entities.MutipleUpload) ([]s3manager.UploadOutput, error)
}
