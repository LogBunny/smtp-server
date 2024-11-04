package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"smtpserver/config"
	"smtpserver/pkg/models"
)

func ForwardMailToBackend(email *models.Email) error {
	data, err := json.Marshal(email)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", config.WEBHOOK_BACKEND_URL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making POST request: %s", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}

	// Print response status and body
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", string(body))
	return nil
}
