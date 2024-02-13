package mssp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const BaseURL = "https://cloud.tenable.com"

type Client struct {
	HTTPClient *http.Client
	ApiKey     string
}

type AccountDetailsV1 struct {
	Email   string `json:"email"`
	Country string `json:"country"`
}

type AddDomainRequest struct {
	ContainerUUID  string `json:"container_uuid"`
	DomainName     string `json:"domain_name"`
	ActivationCode string `json:"activation_code"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		ApiKey:     apiKey,
	}
}

func (c *Client) makeRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, BaseURL+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-ApiKeys", c.ApiKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	return io.ReadAll(resp.Body)
}

// CreateEvaluationAccountV1 creates an evaluation account (V1)
func (c *Client) CreateEvaluationAccountV1(ctx context.Context, accountDetails AccountDetailsV1) ([]byte, error) {
	return c.makeRequest(ctx, "POST", "/mssp/accounts/eval", accountDetails)
}

// ListChildAccounts lists all child accounts
func (c *Client) ListChildAccounts(ctx context.Context) ([]byte, error) {
	return c.makeRequest(ctx, "GET", "/mssp/accounts", nil)
}

// GetChildAccountDetails retrieves details for a specific child account
func (c *Client) GetChildAccountDetails(ctx context.Context, accountUUID string) ([]byte, error) {
	path := fmt.Sprintf("/mssp/accounts/%s", accountUUID)
	return c.makeRequest(ctx, "GET", path, nil)
}

// GetDomainDetails retrieves the details for a specific domain
func (c *Client) GetDomainDetails(ctx context.Context, accountUUID string) ([]byte, error) {
	path := fmt.Sprintf("/mssp/domains/account/%s", accountUUID)
	return c.makeRequest(ctx, "GET", path, nil)
}

// ListDomains returns a list of all child containers and their domains
func (c *Client) ListDomains(ctx context.Context) ([]byte, error) {
	return c.makeRequest(ctx, "GET", "/mssp/domains", nil)
}

// AddDomain adds a new domain to a specific container
func (c *Client) AddDomain(ctx context.Context, request AddDomainRequest) ([]byte, error) {
	return c.makeRequest(ctx, "POST", "/mssp/domains", request)
}

// ListLogos returns a list of logos uploaded to the MSSP Portal
func (c *Client) ListLogos(ctx context.Context) ([]byte, error) {
	return c.makeRequest(ctx, "GET", "/mssp/logos", nil)
}

// Additional structs and methods as needed...
