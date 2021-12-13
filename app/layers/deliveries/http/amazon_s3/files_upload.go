package amazon_s3

import (
	"github.com/chaiyawatkit/ginney"
	"github.com/gin-gonic/gin"
	"upload-file/app/layers/deliveries/http/amazon_s3/models"
	"upload-file/app/utils"
)

func (h *Handler) FilesUpload(c *gin.Context) {
	filesRawData, err := new(models.UploadFilesRequestFile).Parse(c)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	fileMeta, err := h.AmazonS3UseCase.FilesUpload(filesRawData.ToEntity())
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}

	output := new(models.UploadFileResponse).Parse(fileMeta)
	ginney.JSONSuccessResponse(c, output)
}
