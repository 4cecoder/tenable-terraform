package scanner_group

// ScannerGroupUpdateRequest represents the JSON payload for updating a scanner-group group.
type ScannerGroupUpdateRequest struct {
	Name string `json:"name"`
}

// ScannerGroupCreateRequest represents the request payload for creating a scanner-group group.
type ScannerGroupCreateRequest struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}
