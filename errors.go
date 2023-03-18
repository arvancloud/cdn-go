package arvancloud

import (
	"errors"
	"strings"
)

const (
	errEmptyAPIToken = "invalid credentials: API Token must not be empty"

	errMissingDomain      = "required missing domain"
	errMissingDNSRecordID = "required DNS record ID missing"

	errUnmarshalError     = "error unmarshalling the JSON response"
	errUnmarshalErrorBody = "error unmarshalling the JSON response error body"

	errInternalServiceError = "internal service error"
)

var (
	ErrMissingDomain      = errors.New(errMissingDomain)
	ErrMissingDNSRecordID = errors.New(errMissingDNSRecordID)
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

	// Append the main error message to the slice of error messages
	errMessages = append(errMessages, m)

	errors := []string{}

	// Loop through each error in the Errors slice
	// Join them into a comma-separated string
	// Append the string to the errors slice
	for _, e := range e.Errors {
		errors = append(errors, strings.Join(e, ", "))
	}

	// Join the error messages with commas and add it to the errString variable
	errString += strings.Join(errMessages, ", ")

	// If there are errors in the errors slice
	// Join them with a new line and add them to the errString variable
	if len(errors) > 0 {
		errString += "\n" + strings.Join(errors, " \n")
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
