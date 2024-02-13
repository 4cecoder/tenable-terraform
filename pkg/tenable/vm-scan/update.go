package vm_scan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UpdateScanSettings represents the settings to update a scan
type UpdateScanSettings struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Targets     string `json:"text_targets,omitempty"`
}

// UpdateScanPayload represents the payload for updating a scan
type UpdateScanPayload struct {
	Settings UpdateScanSettings `json:"settings"`
}

// SimpleResponse represents a simplified response from the API, for illustration
type SimpleResponse struct {
	Message string `json:"message"`
}

// UpdateScan updates an existing scan with the given settings
func UpdateScan(apiURL, apiKey string, scanID int, payload UpdateScanPayload) (*SimpleResponse, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/scans/%d", apiURL, scanID), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s", apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response SimpleResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here"
//	apiURL := "https://cloud.tenable.com"
//	scanID := 123 // Replace with your actual scan ID
//
//	payload := UpdateScanPayload{
//		Settings: UpdateScanSettings{
//			Name:        "Updated Example Scan",
//			Description: "Updated description",
//			Targets:     "192.168.1.1",
//		},
//	}
//
//	response, err := UpdateScan(apiURL, apiKey, scanID, payload)
//	if err != nil {
//		fmt.Println("Error updating scan:", err)
//		return
//	}
//
//	fmt.Printf("Scan updated successfully: %s\n", response.Message)
//}
