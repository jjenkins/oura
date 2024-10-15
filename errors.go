// Package ouraring provides a client for interacting with the Oura Ring API.
package ouraring

import (
	"fmt"
	"net/http"
)

// APIError represents an error returned by the Oura API.
// It includes the HTTP status code and the error message.
type APIError struct {
	StatusCode int    // HTTP status code returned by the API
	Message    string // Error message returned by the API
}

// Error implements the error interface for APIError.
// It returns a formatted string containing the status code and error message.
func (e *APIError) Error() string {
	return fmt.Sprintf("API error: %d - %s", e.StatusCode, e.Message)
}

// NewAPIError creates a new APIError from an http.Response.
// It extracts the status code and status message from the response.
func NewAPIError(resp *http.Response) *APIError {
	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    resp.Status,
	}
}
