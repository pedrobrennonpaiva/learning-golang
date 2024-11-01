package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Credentials struct
type Credentials struct {
	ClientID     string
	ClientSecret string
	AccountID    string
	SubDomain    string
}

// AuthResponse struct for token response
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// QueryPayload struct for data view query
type QueryPayload struct {
	Query string `json:"query"`
}

func getAccessToken(creds Credentials) (string, error) {
	authURL := fmt.Sprintf("https://%s.auth.marketingcloudapis.com/v2/token", creds.SubDomain)

	payload := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     creds.ClientID,
		"client_secret": creds.ClientSecret,
		"account_id":    creds.AccountID,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", err
	}

	return authResp.AccessToken, nil
}

func getDataViewData(creds Credentials, accessToken, query string) (map[string]interface{}, error) {
	// Endpoint correto para queries SQL
	baseURL := fmt.Sprintf("https://%s.rest.marketingcloudapis.com/v1/sql/query", creds.SubDomain)

	queryPayload := QueryPayload{
		Query: query,
	}

	jsonPayload, err := json.Marshal(queryPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Se a resposta não for 2xx, retorna erro
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error (Status %d): %s", resp.StatusCode, string(body))
	}

	// Parse da resposta JSON
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	// Initialize credentials
	creds := Credentials{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		AccountID:    "YOUR_ACCOUNT_ID",
		SubDomain:    "YOUR_SUBDOMAIN",
	}

	// Get access token
	accessToken, err := getAccessToken(creds)
	if err != nil {
		fmt.Printf("Error getting access token: %v\n", err)
		return
	}

	// Example query for Data View
	query := "SELECT TOP 100 SubscriberKey, EventDate, JobID FROM _Open WHERE EventDate > '2024-01-01'"

	// Get data view data
	results, err := getDataViewData(creds, accessToken, query)
	if err != nil {
		fmt.Printf("Error getting data: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("%+v\n", results)
}
