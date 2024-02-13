package vm_scan

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// ScanList represents the list of scans returned by the API
type ScanList struct {
	Scans []struct {
		ID               int    `json:"id"`
		UUID             string `json:"uuid"`
		Name             string `json:"name"`
		Status           string `json:"status"`
		CreationDate     int64  `json:"creation_date"`
		LastModifiedDate int64  `json:"last_modification_date"`
		// Add other fields as necessary
	} `json:"scans"`
	Folders []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"folders"`
	// Include other response fields as needed
}

// ListScans fetches the list of scans from Tenable
func ListScans(apiURL, apiKey string) (*ScanList, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/scans", apiURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s", apiKey))

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

	var scans ScanList
	err = json.Unmarshal(body, &scans)
	if err != nil {
		return nil, err
	}

	return &scans, nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here"
//	apiURL := "https://cloud.tenable.com"
//
//	scans, err := ListScans(apiURL, apiKey)
//	if err != nil {
//		fmt.Println("Error listing scans:", err)
//		return
//	}
//
//	fmt.Printf("Scans: %+v\n", scans)
//}
