package ouraring

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	token      string
	baseURL    string
}

func NewClient(token string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: time.Second * 10},
		token:      token,
		baseURL:    "https://api.ouraring.com/v2",
	}
}
