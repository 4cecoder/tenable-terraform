package scan_config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// WASConfigUpsertResponse WASConfigDetails represents the detailed information of a WAS scan configuration
type WASConfigUpsertResponse struct {
	ConfigID           string        `json:"config_id"`
	ContainerID        string        `json:"container_id"`
	OwnerID            string        `json:"owner_id"`
	TemplateID         string        `json:"template_id"`
	UserTemplateID     string        `json:"user_template_id,omitempty"`
	Name               string        `json:"name"`
	Targets            []string      `json:"targets"`
	Description        string        `json:"description,omitempty"`
	CreatedAt          string        `json:"created_at"`
	UpdatedAt          string        `json:"updated_at"`
	ScannerID          string        `json:"scanner_id,omitempty"`
	Schedule           interface{}   `json:"schedule,omitempty"`
	DefaultPermissions string        `json:"default_permissions"`
	ResultsVisibility  string        `json:"results_visibility"`
	Permissions        []interface{} `json:"permissions"`
	Notifications      struct {
		Emails []string `json:"emails"`
	} `json:"notifications"`
	// Include settings structure based on the provided JSON example
	Settings struct {
		// Define nested structs for settings as per the JSON example
		// This is a placeholder to indicate where to add the settings structure
	} `json:"settings"`
}

// GetWASConfigDetails fetches details of a WAS scan configuration
func GetWASConfigDetails(apiURL, apiKey string, configID string) (*WASConfigUpsertResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/was/v2/configs/%s", apiURL, configID), nil)
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

	var details WASConfigUpsertResponse
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
//	configID := "your_config_id_here" // Replace with your actual config ID
//
//	details, err := GetWASConfigDetails(apiURL, apiKey, configID)
//	if err != nil {
//		fmt.Println("Error fetching WAS config details:", err)
//		return
//	}
//
//	fmt.Printf("WAS Config Details: %+v\n", details)
//}
