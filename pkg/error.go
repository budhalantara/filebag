package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppError struct {
	Code    int
	Message string
	Errors  []string
}

func NewAppError() *AppError {
	code := http.StatusInternalServerError
	return &AppError{
		Code:    code,
		Message: http.StatusText(code),
	}
}

func NewAppError_BadRequest(message string, errors []string) *AppError {
	code := http.StatusBadRequest
	if message == "" {
		message = http.StatusText(code)
	}

	return &AppError{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

func NewAppError_NotFound(message string) *AppError {
	code := http.StatusNotFound
	if message == "" {
		message = http.StatusText(code)
	}

	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewAppError_UnprocessableEntity(message string, errors []string) *AppError {
	code := http.StatusUnprocessableEntity
	if message == "" {
		message = http.StatusText(code)
	}

	return &AppError{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}

func (ae AppError) ToApiResponse(c echo.Context) error {
	return c.JSON(ae.Code, ApiResponse{
		Message: ae.Message,
		Errors:  ae.Errors,
	})
}
