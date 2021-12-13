package amazon_s3

import (
	"github.com/gin-gonic/gin"
	"upload-file/app/layers/usecases/amazon_s3"
)

type ResponseError struct {
	Message string `json:"message"`
}

type Handler struct {
	AmazonS3UseCase amazon_s3.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, amazon amazon_s3.UseCase) {
	handler := &Handler{
		AmazonS3UseCase: amazon,
	}
	v1 := ginEngine.Group("v1")
	{
		v1.POST("/files-upload", handler.FilesUpload)
		v1.POST("/multiple-uploads", handler.MultipleFilesUpload)
	}
}
