package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Response represents the standard API response format
type Response struct {
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}, message string) Response {
	return Response{
		Message: message,
		Success: true,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string) Response {
	return Response{
		Message: message,
		Success: false,
	}
}

// JSONResponse sends a JSON response with the given status code
func JSONResponse(c *gin.Context, status int, response Response) {
	c.JSON(status, response)
}

// Success sends a success response with HTTP 200 status
func Success(c *gin.Context, data interface{}, message string) {
	JSONResponse(c, http.StatusOK, NewSuccessResponse(data, message))
}

// Error sends an error response with HTTP 400 status
func Error(c *gin.Context, message string) {
	JSONResponse(c, http.StatusBadRequest, NewErrorResponse(message))
}

// ErrorWithStatus sends an error response with custom status code
func ErrorWithStatus(c *gin.Context, status int, message string) {
	JSONResponse(c, status, NewErrorResponse(message))
}
