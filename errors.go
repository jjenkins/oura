package ouraring

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: %d - %s", e.StatusCode, e.Message)
}

func NewAPIError(resp *http.Response) *APIError {
	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    resp.Status,
	}
}
