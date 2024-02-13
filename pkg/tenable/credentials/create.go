package credentials

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://cloud.tenable.com"

// CreateCredential creates a new managed credential in Tenable.io.
func CreateCredential(apiKey string, credential Credential) (*http.Response, error) {
	client := &http.Client{}
	endpoint := baseURL + "/credentials"

	// Marshal the credential struct to JSON
	credentialJSON, err := json.Marshal(credential)
	if err != nil {
		return nil, err
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(credentialJSON))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ApiKeys", apiKey) // Adjust header as needed based on actual API key requirements

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Optional: Handle non-200 responses
	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		var errorResp ErrorResponse
		if err := json.Unmarshal(body, &errorResp); err == nil {
			return nil, fmt.Errorf("API error: %s", errorResp.Message)
		}
		return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	return resp, nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here" // Replace with your actual API key
//	credential := Credential{
//		Name:        "Example Credential",
//		Description: "An example credential for demonstration",
//		Type:        "ssh", // For example, adjust based on actual credential type needed
//		Settings: map[string]interface{}{
//			"username": "user",
//			"password": "pass",
//
//		},
//
//		Permissions: []Permission{
//			// Define permissions as needed
//		},
//	}
//
//	response, err := CreateCredential(apiKey, credential)
//	if err != nil {
//		fmt.Println("Error creating credential:", err)
//		return
//	}
//
//	fmt.Printf("Credential created successfully, response status: %s\n", response.Status)
//}
