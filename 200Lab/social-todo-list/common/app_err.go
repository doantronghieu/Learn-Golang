package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Represents a custom error type with additional fields for status code,
// root error, error message, log information, and a key.
type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullErrorResponse(statusCode int, rootErr error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(rootErr error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(rootErr error, message, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    rootErr,
		Message:    message,
		Key:        key,
	}
}

// Recursively retrieves the root error in case the current error is nested.
func (e *AppError) RootError() error {
	// Check if the root error is also an AppError, and recursively call RootError
	// to find the ultimate root error.
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	// Return the ultimate root error if the current root error is not an AppError.
	return e.RootErr
}

// Returns the string representation of the error by calling Error() on the ultimate root error.
func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCustomError(rootErr error, message string, key string) *AppError {
	if rootErr != nil {
		return NewErrorResponse(rootErr, message, rootErr.Error(), key)
	}
	return NewErrorResponse(errors.New(message), message, message, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(
		err, "invalid request", err.Error(), "ErrInvalidRequest",
	)
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong in the server",
		err.Error(),
		"ErrInternal",
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotFoundEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Found %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotFound%s", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		"You have no permission",
		"ErrNoPermission",
	)
}

var RecordNotFound = errors.New("record not found")
