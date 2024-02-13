package vm_scan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// ScanSettings represents the settings for the scan
type ScanSettings struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Targets     string `json:"targets"`
}

// ScanPayload represents the payload for creating a scan
type ScanPayload struct {
	Settings ScanSettings `json:"settings"`
}

// ScanResponse represents the response from the scan creation API
type ScanResponse struct {
	ScanID int `json:"scan_id"`
}

// CreateScan creates a new scan with the given settings
func CreateScan(apiURL, apiKey string, payload ScanPayload) (*ScanResponse, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL+"/scans", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ApiKeys", "accessKey="+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var scanResp ScanResponse
	err = json.Unmarshal(body, &scanResp)
	if err != nil {
		return nil, err
	}

	return &scanResp, nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here"
//	apiURL := "https://cloud.tenable.com"
//
//	payload := ScanPayload{
//		Settings: ScanSettings{
//			Name:        "Example Scan",
//			Description: "This is an example scan",
//			Targets:     "127.0.0.1",
//		},
//	}
//
//	scanResp, err := CreateScan(apiURL, apiKey, payload)
//	if err != nil {
//		fmt.Println("Error creating scan:", err)
//		return
//	}
//
//	fmt.Printf("Scan created successfully: %+v\n", scanResp)
//}
