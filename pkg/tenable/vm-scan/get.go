package vm_scan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ScanDetails represents the detailed information of a scan
type ScanDetails struct {
	Info struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Status      string `json:"status"`
		ScanStart   string `json:"scan_start"`
		ScanEnd     string `json:"scan_end"`
		ScanType    string `json:"scan_type"`
		Targets     string `json:"targets"`
		ObjectID    string `json:"object_id"`
		Owner       string `json:"owner"`
		Policy      string `json:"policy"`
		ScannerName string `json:"scanner_name"`
		// Add other fields as needed
	} `json:"info"`
	// Include other top-level response structures as needed
}

// GetScanDetails fetches details for a specific scan from Tenable
func GetScanDetails(apiURL string, apiKey string, scanID int) (*ScanDetails, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/scans/%d", apiURL, scanID), nil)
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

	var details ScanDetails
	err = json.Unmarshal(body, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here"
//	apiURL := "https://cloud.tenable.com"
//	scanID := 123 // Replace with your actual scan ID
//
//	details, err := GetScanDetails(apiURL, apiKey, scanID)
//	if err != nil {
//		fmt.Println("Error fetching scan details:", err)
//		return
//	}
//
//	fmt.Printf("Scan Details: %+v\n", details)
//}
//
