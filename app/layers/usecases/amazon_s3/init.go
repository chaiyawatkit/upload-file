package amazon_s3

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"upload-file/app/entities"
	"upload-file/app/layers/repositories/amazon_s3"
)

type useCase struct {
	AmazonRepo amazon_s3.Repo
}

func Init(amazonRepo amazon_s3.Repo) UseCase {
	return &useCase{
		AmazonRepo: amazonRepo,
	}
}

type UseCase interface {
	FilesUpload(upload entities.FilesUpload) (*entities.S3Response, error)
	MultipleFilesUpload(upload entities.MutipleUpload) ([]s3manager.UploadOutput, error)
}
