package scan_config

import (
	"fmt"
	"net/http"
)

// DeleteWASConfig deletes a web application scanning configuration by its ID
func DeleteWASConfig(apiURL, apiKey, configID string) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/was/v2/configs/%s", apiURL, configID), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ApiKeys", fmt.Sprintf("accessKey=%s", apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code for success (200 OK or 204 No Content typically)
	if resp.StatusCode != 200 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete WAS config, status code: %d", resp.StatusCode)
	}

	return nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here"
//	apiURL := "https://cloud.tenable.com"
//	configID := "your_config_id_here" // Replace with the actual config ID you want to delete
//
//	err := DeleteWASConfig(apiURL, apiKey, configID)
//	if err != nil {
//		fmt.Println("Error deleting WAS config:", err)
//	} else {
//		fmt.Println("WAS config deleted successfully")
//	}
//}
