package amazon_s3

import (
	"upload-file/app/entities"
	"upload-file/app/errors"
)

func (u *useCase) FilesUpload(upload entities.FilesUpload) (*entities.S3Response, error) {
	fileMeta, err := u.AmazonRepo.FilesUpload(upload)
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	return fileMeta, nil
}
