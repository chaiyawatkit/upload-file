package amazon_s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"upload-file/app/entities"
	"upload-file/app/env"

	"time"
)

const (
	AwsS3Region = "ap-southeast-1"
)

func (r *repo) FilesUpload(upload entities.FilesUpload) (*entities.S3Response, error) {
	env.Init()
	now := time.Now()
	sess, err := connectAWS()
	filename := fmt.Sprintf("%d-%v", now.UnixNano(), upload.FilesHeader.Filename)
	uploader := s3manager.NewUploader(sess)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(env.AwsS3Bucket),
		Key:    aws.String(filename),
		Body:   upload.Files,
	})
	if err != nil {
		return nil, err
	}
	resDataS3 := &entities.S3Response{
		Location: output.Location,
	}
	return resDataS3, nil
}

func connectAWS() (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(AwsS3Region),
			Credentials: credentials.NewStaticCredentials(env.S3AKID, env.S3SecretKey, ""),
		},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
