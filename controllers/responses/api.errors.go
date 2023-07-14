package responses

import (
	"github.com/gin-gonic/gin"
)

// APIError is the error format returned by the API
// swagger:model APIError
type APIError struct {
	// Message of the error
	// example: Internal server error
	Message string `json:"message"`
	// Code of the error
	// example: internal_server_error
	Code string `json:"code"`
	// Status of the error
	// example: 500
	Status int `json:"status"`
	// Context of the error
	// example: {"body": "body of the request"}
	Context map[string]interface{} `json:"context"`
	// OriginalError is the original error
	// example: Error parsing JSON
	OriginalError string `json:"error"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError() *APIError {
	return &APIError{
		Message:       "Internal server error",
		Code:          "internal_server_error",
		Status:        500,
		Context:       nil,
		OriginalError: "",
	}
}

func (e *APIError) WithMessage(message string) *APIError {
	e.Message = message
	return e
}

func (e *APIError) WithCode(code string) *APIError {
	e.Code = code
	return e
}

func (e *APIError) WithStatus(status int) *APIError {
	e.Status = status
	return e
}

func (e *APIError) WithContext(context map[string]interface{}) *APIError {
	e.Context = context
	return e
}

func (e *APIError) WithError(err error) *APIError {
	e.OriginalError = err.Error()
	return e
}

func (e *APIError) ProduceAPIError(c *gin.Context) {
	c.JSON(e.Status, e)
}
