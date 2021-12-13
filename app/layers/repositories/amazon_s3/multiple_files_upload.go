package amazon_s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"upload-file/app/entities"
	"upload-file/app/env"

	"math/rand"
	"mime/multipart"
	"path/filepath"
	"time"
)

func (r *repo) MultipleFilesUpload(upload entities.MutipleUpload) ([]s3manager.UploadOutput, error) {
	env.Init()
	now := time.Now()
	sess, err := connectAWS()
	var outputArr []s3manager.UploadOutput
	var myFiles multipart.File
	for i := 0; i < len(upload.FilesHeader); i++ {
		myFiles, err = upload.FilesHeader[i].Open()
		filename := fmt.Sprintf("%d-%v%v", now.UnixNano(), rand.Float64(), filepath.Ext(upload.FilesHeader[i].Filename))
		uploader := s3manager.NewUploader(sess)
		output, errOutput := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(env.AwsS3Bucket),
			Key:    aws.String(filename),
			Body:   myFiles,
		})
		myFiles.Close()
		if errOutput != nil {
			break
			return nil, err
		}
		outputArr = append(outputArr, *output)

	}
	return outputArr, nil
}
