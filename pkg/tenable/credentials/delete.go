package credentials

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteCredential deletes a managed credential by its UUID.
func DeleteCredential(apiKey, uuid string) error {
	client := &http.Client{}
	endpoint := fmt.Sprintf("https://cloud.tenable.com/credentials/%s", uuid)

	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	// Set necessary headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-ApiKeys", apiKey) // Adjust header as needed based on actual API key requirements

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Check for non-200 response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here" // Replace with your actual API key
//	uuid := "credential_uuid_here" // Replace with the UUID of the credential you want to delete
//
//	err := DeleteCredential(apiKey, uuid)
//	if err != nil {
//		fmt.Println("Error deleting credential:", err)
//	} else {
//		fmt.Println("Credential deleted successfully")
//	}
//}
