package downloads

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client struct represents the Downloads API client
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Token      string
}

// NewClient creates a new instance of the Downloads API client
func NewClient(token string) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		BaseURL:    "https://www.tenable.com/downloads/api/v2",
		Token:      token,
	}
}

// makeRequest is a helper function to make an HTTP request to the Downloads API
func (c *Client) makeRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// ListProductPages lists all the product pages available in the Downloads API
func (c *Client) ListProductPages(ctx context.Context) ([]byte, error) {
	return c.makeRequest(ctx, "GET", "/pages", nil)
}

// ListDownloadsForProduct lists all downloadable files for a given product
func (c *Client) ListDownloadsForProduct(ctx context.Context, slug string) ([]byte, error) {
	return c.makeRequest(ctx, "GET", fmt.Sprintf("/pages/%s", slug), nil)
}

// DownloadFile downloads a specific file for a given product
func (c *Client) DownloadFile(ctx context.Context, slug, fileName string) ([]byte, error) {
	return c.makeRequest(ctx, "GET", fmt.Sprintf("/pages/%s/files/%s", slug, fileName), nil)
}
