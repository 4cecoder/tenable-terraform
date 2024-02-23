package main

import (
	"github.com/4cecoder/tenable-terraform/pkg/tenable/credentials"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCredentialHandler(c *gin.Context) {
	// Extract the credential ID from the URL parameters
	id := c.Param("id")

	apiKey := "your_api_key" // Replace with your actual API key
	credentialDetail, err := credentials.GetCredentialDetail(apiKey, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send the credential details back to the client
	c.JSON(http.StatusOK, credentialDetail)
}

func UpdateCredentialHandler(c *gin.Context) {
	// Extract the credential ID from the URL parameters
	id := c.Param("id")

	// Extract the updated credential details from the request body
	var update credentials.CredentialUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiKey := "your_api_key" // Replace with your actual API key
	err := credentials.UpdateCredential(apiKey, id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send a success response back to the client
	c.JSON(http.StatusOK, gin.H{"status": "Credential updated successfully"})
}

func CreateCredentialHandler(c *gin.Context) {
	var credential credentials.Credential
	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiKey := "your_api_key" // Replace with your actual API key
	response, err := credentials.CreateCredential(apiKey, credential)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": response.Status})
}

func DeleteCredentialHandler(c *gin.Context) {
	// Extract the credential ID from the URL parameters
	id := c.Param("id")

	apiKey := "your_api_key" // Replace with your actual API key
	err := credentials.DeleteCredential(apiKey, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send a success response back to the client
	c.JSON(http.StatusOK, gin.H{"status": "Credential deleted successfully"})
}

func ListCredentialsHandler(c *gin.Context) {
	apiKey := "your_api_key" // Replace with your actual API key
	credentialList, err := credentials.ListCredentials(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send the list of credentials back to the client
	c.JSON(http.StatusOK, credentialList)
}
