package scanner_group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateScannerGroup sends a request to create a new scanner-group group in Tenable.io.
func CreateScannerGroup(apiKey, name, groupType, description string) error {
	url := "https://cloud.tenable.com/scanner-groups"
	payload := ScannerGroupCreateRequest{
		Name:        name,
		Type:        groupType,
		Description: description,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-ApiKeys", fmt.Sprintf("accessKey=%s", apiKey)) // Adjust based on actual API key header format

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println("Scanner group created successfully")
	return nil
}

//func main() {
//	apiKey := "your_api_key_here" // Replace with your actual API key
//	name := "New Scanner Group"
//	groupType := "type_here" // Set the appropriate type
//	description := "Description of the new scanner-group group"
//
//	err := CreateScannerGroup(apiKey, name, groupType, description)
//	if err != nil {
//		fmt.Println("Error creating scanner-group group:", err)
//	}
//}
