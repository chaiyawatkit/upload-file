package amazon_s3

import (
	"github.com/chaiyawatkit/ginney"
	"github.com/gin-gonic/gin"
	"upload-file/app/layers/deliveries/http/amazon_s3/models"
	"upload-file/app/utils"
)

func (h *Handler) MultipleFilesUpload(c *gin.Context) {

	filesRawData, err := new(models.UploadMultipleFilesRequestFile).Parse(c)
	if err != nil {

		utils.JSONErrorResponse(c, err)
		return
	}
	dataFile, err := filesRawData.IsValid(c)

	if err != nil {

		utils.JSONErrorResponse(c, err)
		return
	}

	fileMeta, err := h.AmazonS3UseCase.MultipleFilesUpload(dataFile.ToEntity())
	if err != nil {

		utils.JSONErrorResponse(c, err)
		return
	}
	output := new(models.MultipleFileResponseJSON).ResponseJSON(fileMeta)
	ginney.JSONSuccessResponse(c, output)
}
