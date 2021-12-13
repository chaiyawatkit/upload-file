package healthcheck

import (
	"github.com/gin-gonic/gin"
	"upload-file/app/utils"
)

func (handler *handler) Health(c *gin.Context) {
	utils.JSONSuccessResponse(c, nil)
}
