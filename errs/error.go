package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

//helper func
func NewNotFoundError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusInternalServerError}
}
