package scan_config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// WASConfigPayload represents the payload for creating a WAS scan configuration
type WASConfigPayload struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// Include other necessary fields based on the API documentation
}

// WASConfigResponse represents the response from creating a WAS scan configuration
type WASConfigResponse struct {
	ConfigID string `json:"config_id"`
	// Adjust based on the actual API response structure
}

// CreateWASConfig creates a new WAS scan configuration
func CreateWASConfig(apiURL, apiKey string, payload WASConfigPayload) (*WASConfigResponse, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/was/v2/configs", apiURL), bytes.NewBuffer(jsonPayload))
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

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var configResp WASConfigResponse
	err = json.Unmarshal(body, &configResp)
	if err != nil {
		return nil, err
	}

	return &configResp, nil
}

// Example Usage
//func main() {
//	apiKey := "your_api_key_here"
//	apiURL := "https://cloud.tenable.com"
//
//	payload := WASConfigPayload{
//		Name:        "New WAS Config",
//		Description: "Description of the new WAS config",
//		// Populate other necessary fields
//	}
//
//	configResp, err := CreateWASConfig(apiURL, apiKey, payload)
//	if err != nil {
//		fmt.Println("Error creating WAS config:", err)
//		return
//	}
//
//	fmt.Printf("WAS Config created successfully: %+v\n", configResp)
//}
//
