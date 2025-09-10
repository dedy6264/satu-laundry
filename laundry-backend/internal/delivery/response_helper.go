package delivery

import (
	"laundry-backend/internal/entities"

	"github.com/labstack/echo/v4"
)

// SuccessResponse generates a success response with data
func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	response := entities.APIResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	return c.JSON(statusCode, response)
}

// ErrorResponse generates an error response
func ErrorResponse(c echo.Context, statusCode int, message string, err string) error {
	response := entities.APIResponse{
		Status:  statusCode,
		Message: message,
		Error:   err,
	}
	return c.JSON(statusCode, response)
}

// MessageResponse generates a response with only a message (no data)
func MessageResponse(c echo.Context, statusCode int, message string) error {
	response := entities.APIResponse{
		Status:  statusCode,
		Message: message,
	}
	return c.JSON(statusCode, response)
}