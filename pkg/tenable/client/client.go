package client

import (
	"bytes"
	"fmt"
	"net/http"
)

// Client holds information necessary to make requests to Tenable.io
type Client struct {
	AccessKey string
	SecretKey string
	BaseURL   string
}

// NewClient returns a new Tenable.io API client
func NewClient(accessKey, secretKey string) *Client {
	return &Client{
		AccessKey: accessKey,
		SecretKey: secretKey,
		BaseURL:   "https://cloud.tenable.com",
	}
}

// makeRequest prepares a new API request with necessary headers
func (c *Client) makeRequest(method, endpoint string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL+endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s; secretKey=%s", c.AccessKey, c.SecretKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}
