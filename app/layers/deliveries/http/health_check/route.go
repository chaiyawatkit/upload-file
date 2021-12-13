package healthcheck

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
}

// NewEndpointHTTPHandler setup router
func NewEndpointHTTPHandler(ginEngine *gin.Engine) {
	handler := handler{}
	ginEngine.GET("/health", handler.Health)
}
