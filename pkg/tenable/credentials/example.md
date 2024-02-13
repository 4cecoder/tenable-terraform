


```go 
func main() {
    apiKey := "your_api_key_here" // Replace with your actual API key
    credential := Credential{
    Name:        "Example Credential",
    Description: "An example credential for demonstration",
    Type:        "ssh", // For example, adjust based on actual credential type needed
    Settings: map[string]interface{}{
    "username": "user",
    "password": "pass",
	
    },
	
Permissions: []Permission{
// Define permissions as needed
},
}

    response, err := CreateCredential(apiKey, credential)
    if err != nil {
    fmt.Println("Error creating credential:", err)
    return
}

    fmt.Printf("Credential created successfully, response status: %s\n", response.Status)
}

```
