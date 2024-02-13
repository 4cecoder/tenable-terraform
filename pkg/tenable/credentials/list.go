package credentials

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ListCredentials fetches a list of managed credentials with optional filtering and pagination.
func ListCredentials(apiKey string, options ...func(*url.Values)) (*CredentialList, error) {
	client := &http.Client{}
	baseURL := "https://cloud.tenable.com/credentials"
	queryParams := url.Values{}

	// Apply any provided query parameters
	for _, option := range options {
		option(&queryParams)
	}

	// Construct the full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// Create a new GET request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-ApiKeys", apiKey) // Adjust this header based on actual API key requirements

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
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	// Parse the JSON response
	var credentialList CredentialList
	if err := json.NewDecoder(resp.Body).Decode(&credentialList); err != nil {
		return nil, err
	}

	return &credentialList, nil
}

//func main() {
//	apiKey := "your_api_key_here" // Replace with your actual API key
//
//	// Example options for filtering and pagination
//	filterByName := func(values *url.Values) {
//		values.Add("f", "name:match:example")
//	}
//	sortByNameAsc := func(values *url.Values) {
//		values.Add("sort", "name:asc")
//	}
//	limitResults := func(values *url.Values) {
//		values.Add("limit", "10")
//	}
//
//	credentials, err := ListCredentials(apiKey, filterByName, sortByNameAsc, limitResults)
//	if err != nil {
//		fmt.Println("Error listing credentials:", err)
//		return
//	}
//
//	fmt.Printf("Total Credentials Found: %d\n", credentials.Total)
//	for _, credential := range credentials.Credentials {
//		fmt.Printf("ID: %s, Name: %s\n", credential.ID, credential.Name)
//	}
//}
