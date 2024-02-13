package scanner_group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// UpdateScannerGroup updates the specified scanner-group group's name.
func UpdateScannerGroup(apiKey, groupID, newName string) error {
	url := fmt.Sprintf("https://cloud.tenable.com/scanner-groups/%s", groupID)
	payload := ScannerGroupUpdateRequest{Name: newName}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-ApiKeys", apiKey) // Adjust the header based on actual API key requirements

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update scanner-group group with status code: %d, response: %s", resp.StatusCode, string(body))
	}

	fmt.Println("Scanner group updated successfully")
	return nil
}

//func main() {
//	apiKey := "your_api_key_here" // Replace with your actual API key
//	groupID := "group_id_here"    // Replace with the actual group ID
//	newName := "New Scanner Group Name"
//
//	if err := UpdateScannerGroup(apiKey, groupID, newName); err != nil {
//		fmt.Println("Error updating scanner-group group:", err)
//	}
//}
