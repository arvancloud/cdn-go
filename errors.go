package arvancloud

import (
	"errors"
	"strings"
)

const (
	errEmptyAPIToken        = "invalid credentials: API Token must not be empty"
	errMissingDomain        = "required missing domain"
	errUnmarshalError       = "error unmarshalling the JSON response"
	errInternalServiceError = "internal service error"
	errUnmarshalErrorBody   = "error unmarshalling the JSON response error body"
)

var (
	ErrMissingDomain = errors.New(errMissingDomain)
)

type ErrorType string

const (
	ErrorTypeRequest        ErrorType = "request"
	ErrorTypeAuthentication ErrorType = "authentication"
	ErrorTypeAuthorization  ErrorType = "authorization"
	ErrorTypeNotFound       ErrorType = "not_found"
	ErrorTypeRateLimit      ErrorType = "rate_limit"
	ErrorTypeService        ErrorType = "service"
)

type Error struct {
	// The classification of error encountered.
	Type ErrorType

	// StatusCode is the HTTP status code from the response.
	Code int

	// Message is the error message.
	Message string

	// Errors is a list of all the response error.
	Errors map[string][]string
}

// Error will return a human readable error message.
func (e Error) Error() string {
	var errString string
	errMessages := []string{}
	m := ""
	if e.Message != "" {
		m += e.Message
	}

	errMessages = append(errMessages, m)

	errors := []string{}
	for _, e := range e.Errors {
		errors = append(errors, strings.Join(e, ", "))
	}

	errString += strings.Join(errMessages, ", ")

	if len(errors) > 0 {
		errString += "\n" + strings.Join(errors, "  \n")
	}

	return errString
}

// RequestError is for 4xx errors.
type RequestError struct {
	arvancloudError *Error
}

func (e RequestError) Error() string {
	return e.arvancloudError.Error()
}

func (e RequestError) Errors() map[string][]string {
	return e.arvancloudError.Errors
}

func (e RequestError) ErrorCode() int {
	return e.arvancloudError.Code
}

func (e RequestError) ErrorMessage() string {
	return e.arvancloudError.Message
}

func (e RequestError) Type() ErrorType {
	return e.arvancloudError.Type
}

func NewRequestError(e *Error) RequestError {
	return RequestError{
		arvancloudError: e,
	}
}

// ServiceError is for 5xx errors.
type ServiceError struct {
	arvancloudError *Error
}

func (e ServiceError) Error() string {
	return e.arvancloudError.Error()
}

func (e ServiceError) Errors() map[string][]string {
	return e.arvancloudError.Errors
}

func (e ServiceError) ErrorCode() int {
	return e.arvancloudError.Code
}

func (e ServiceError) ErrorMessage() string {
	return e.arvancloudError.Message
}

func (e ServiceError) Type() ErrorType {
	return e.arvancloudError.Type
}

func NewServiceError(e *Error) ServiceError {
	return ServiceError{
		arvancloudError: e,
	}
}
