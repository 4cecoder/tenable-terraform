package credentials

// CredentialDetail represents the details of a managed credential returned by Tenable.io.
type CredentialDetail struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Settings    map[string]interface{} `json:"settings"`
	Permissions []Permission           `json:"permissions,omitempty"`
	// Add other fields as per the API response
}

// Credential represents the request body for creating a new credential.
type Credential struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Settings    map[string]interface{} `json:"settings"`
	Permissions []Permission           `json:"permissions,omitempty"`
}

// CredentialList represents a list of managed credentials.
type CredentialList struct {
	Credentials []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Type        string `json:"type"`
		// Include other fields as necessary.
	} `json:"credentials"`
	// Assuming the response includes some form of pagination or total count
	Total int `json:"total"`
}

// CredentialUpdate represents the body for updating a managed credential.
type CredentialUpdate struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	AdHoc       *bool                  `json:"ad_hoc,omitempty"`
	Settings    map[string]interface{} `json:"settings,omitempty"`
	Permissions []Permission           `json:"permissions,omitempty"`
}

// Permission represents user permissions for the managed credential.
type Permission struct {
	Type      string `json:"type"`
	Value     int    `json:"value"`
	GranteeID string `json:"grantee_uuid,omitempty"` // Adjust as needed based on actual API requirements
}

// ErrorResponse represents a typical error response from the API.
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
