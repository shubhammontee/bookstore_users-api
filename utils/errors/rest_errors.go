package errors

import (
	"errors"
	"net/http"
)

type (
	//RestErr ...
	RestErr struct {
		Message string `json:"message,omitempty"`
		Status  int    `json:"status,omitempty"`
		Error   string `json:"error,omitempty"`
	}
)

//NewError ...
func NewError(msg string) error {
	return errors.New(msg)
}

//NewBadRequestError ...
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError ...
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

//NewInternalServerError ...
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}
