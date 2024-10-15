// Package ouraring provides a client for interacting with the Oura Ring API.
package ouraring

import (
	"net/http"
	"time"
)

// Client represents an Oura API client. It handles authentication and
// provides methods for making API requests.
type Client struct {
	httpClient *http.Client // HTTP client used for making API requests
	token      string       // Authentication token for the Oura API
	baseURL    string       // Base URL for the Oura API
}

// NewClient creates and returns a new Oura API client.
// It takes an authentication token as a parameter and sets up the client
// with default configuration.
//
// The client is configured with a 10-second timeout for all requests.
// If you need a different timeout, you can modify the httpClient after creation.
func NewClient(token string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: time.Second * 10},
		token:      token,
		baseURL:    "https://api.ouraring.com/v2",
	}
}
