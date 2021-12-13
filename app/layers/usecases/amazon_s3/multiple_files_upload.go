package amazon_s3

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"upload-file/app/entities"
	"upload-file/app/errors"
)

func (u *useCase) MultipleFilesUpload(upload entities.MutipleUpload) ([]s3manager.UploadOutput, error) {
	fileMeta, err := u.AmazonRepo.MultipleFilesUpload(upload)
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	return fileMeta, nil

}
