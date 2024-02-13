package credentials

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetCredentialDetail fetches the details of a managed credential by its UUID.
func GetCredentialDetail(apiKey, uuid string) (*CredentialDetail, error) {
	client := &http.Client{}
	endpoint := fmt.Sprintf("%s/credentials/%s", baseURL, uuid)

	// Create a new GET request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-ApiKeys", apiKey) // Adjust header as needed based on actual API key requirements

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Check for non-200 response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Parse the JSON response into the CredentialDetail struct
	var credentialDetail CredentialDetail
	if err := json.NewDecoder(resp.Body).Decode(&credentialDetail); err != nil {
		return nil, err
	}

	return &credentialDetail, nil
}
