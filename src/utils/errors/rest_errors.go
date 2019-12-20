package errors

import (
	"errors"
	"net/http"
)

type RestErrors struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string) error {
	return errors.New(message)
}

func BadRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad  request",
	}
}

func NotFoundError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}
}

func InternalServerError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
}
