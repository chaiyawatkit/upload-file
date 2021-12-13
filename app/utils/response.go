package utils

import (
	"github.com/gin-gonic/gin"
	"upload-file/app/errors"

	"net/http"
)

const (
	// StatusFail fail message
	StatusFail string = "fail"

	// StatusSuccess success message
	StatusSuccess string = "success"
)

const (
	// MessageOk OK message
	MessageOk string = "OK"
)

type baseSuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"OK"`
	Data    interface{} `json:"data,omitempty"`
}

type successResponseWithPagination struct {
	Status   string      `json:"status" example:"success"`
	Message  string      `json:"message" example:"OK"`
	Metadata interface{} `json:"_metadata"`
	Data     interface{} `json:"data"`
}

// JSONSuccessResponse success response without meta data
func JSONSuccessResponse(c *gin.Context, data interface{}) {
	r := new(baseSuccessResponse)
	r.Status = StatusSuccess
	r.Message = MessageOk
	r.Data = data
	c.AbortWithStatusJSON(http.StatusOK, *r)
}

// JSONSuccessResponseWithPagination success response with meta data
func JSONSuccessResponseWithPagination(c *gin.Context, data interface{}, meta interface{}) {
	r := new(successResponseWithPagination)
	r.Status = StatusSuccess
	r.Message = MessageOk
	r.Metadata = meta
	r.Data = data
	c.AbortWithStatusJSON(http.StatusOK, *r)
}

// errorResponse base error response
type errorResponse struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message" example:"Error message will be show here"`
}

// JSONErrorResponse response error json
func JSONErrorResponse(c *gin.Context, err error) {
	statusCode := http.StatusInternalServerError
	message := ""

	switch err.(type) {
	case errors.ParameterError:
		statusCode = http.StatusBadRequest
		message = err.(errors.ParameterError).Error()
	case errors.UnprocessableEntity:
		statusCode = http.StatusBadRequest
		message = err.(errors.UnprocessableEntity).Error()
	case errors.InternalError:
		statusCode = http.StatusInternalServerError
		message = err.(errors.InternalError).Error()
	case errors.RecordNotFoundError:
		statusCode = http.StatusNotFound
		message = err.(errors.RecordNotFoundError).Error()
		err = nil
	default:
		message = err.Error()
	}

	errorResponse := errorResponse{
		Status:  StatusFail,
		Message: message,
	}

	c.AbortWithStatusJSON(statusCode, errorResponse)
}
