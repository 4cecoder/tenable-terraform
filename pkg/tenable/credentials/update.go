package credentials

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// UpdateCredential updates the details of a managed credential by its UUID.
func UpdateCredential(apiKey, uuid string, update CredentialUpdate) error {
	client := &http.Client{}
	endpoint := fmt.Sprintf("https://cloud.tenable.com/credentials/%s", uuid)

	// Marshal the update struct to JSON
	updateJSON, err := json.Marshal(update)
	if err != nil {
		return err
	}

	// Create a new PUT request
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(updateJSON))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
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

	// Optional: Handle non-200 responses
	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		var errorResp ErrorResponse
		if err := json.Unmarshal(body, &errorResp); err == nil {
			return fmt.Errorf("API error: %s", errorResp.Message)
		}
		return fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	return nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here" // Replace with your actual API key
//	uuid := "credential_uuid_here" // Replace with the UUID of the credential
//
//	update := CredentialUpdate{
//		Name:        "Updated Credential Name",
//		Description: "Updated description",
//		// Set other fields as needed
//	}
//
//	err := UpdateCredential(apiKey, uuid, update)
//	if err != nil {
//		fmt.Println("Error updating credential:", err)
//	} else {
//		fmt.Println("Credential updated successfully")
//	}
//}
