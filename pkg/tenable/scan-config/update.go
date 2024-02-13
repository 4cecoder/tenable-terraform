package scan_config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ScanConfigUpdatePayload represents the payload for updating a scan configuration
type ScanConfigUpdatePayload struct {
	// Define your payload structure here according to the API documentation
	// This example assumes generic fields for demonstration purposes
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// Add other fields as per the Tenable API documentation
}

// ScanConfigUpdateResponse represents the response from updating a scan configuration
type ScanConfigUpdateResponse struct {
	// Define your response structure here
	// This example assumes a generic response for demonstration purposes
	Message string `json:"message"`
}

// UpdateScanConfig sends a request to update an existing scan configuration
func UpdateScanConfig(apiURL, apiKey string, configID string, payload ScanConfigUpdatePayload) (*ScanConfigUpdateResponse, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	requestURL := fmt.Sprintf("%s/scans/%s", apiURL, configID)
	req, err := http.NewRequest("PUT", requestURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		return nil, fmt.Errorf("API request error: %s", string(body))
	}

	var response ScanConfigUpdateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here"         // Replace with your actual API key
//	apiURL := "https://cloud.tenable.com" // Base URL for the Tenable API
//	configID := "your_config_id_here"     // Replace with the ID of the scan configuration you wish to update
//
//	payload := ScanConfigUpdatePayload{
//		Name:        "Updated Scan Name",
//		Description: "Updated description here",
//		// Populate other fields as needed
//	}
//
//	response, err := UpdateScanConfig(apiURL, apiKey, configID, payload)
//	if err != nil {
//		fmt.Printf("Error updating scan config: %v\n", err)
//		return
//	}
//
//	fmt.Printf("Scan config updated successfully: %+v\n", response)
//}
