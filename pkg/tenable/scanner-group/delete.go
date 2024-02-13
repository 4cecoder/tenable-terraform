package scanner_group

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteScannerGroup sends a request to delete a scanner-group group by its ID.
func DeleteScannerGroup(apiKey, groupID string) error {
	client := &http.Client{}
	url := fmt.Sprintf("https://cloud.tenable.com/scanner-groups/%s", groupID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-ApiKeys", fmt.Sprintf("accessKey=%s", apiKey)) // Adjust based on actual API key header format

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete scanner-group group with status code: %d", resp.StatusCode)
	}

	fmt.Println("Scanner group deleted successfully")
	return nil
}

//func main() {
//	apiKey := "your_api_key_here" // Replace with your actual API key
//	groupID := "group_id_here"    // Replace with the scanner-group group ID you want to delete
//
//	err := DeleteScannerGroup(apiKey, groupID)
//	if err != nil {
//		fmt.Println("Error deleting scanner-group group:", err)
//	} else {
//		fmt.Println("Scanner group deleted successfully")
//	}
//}
