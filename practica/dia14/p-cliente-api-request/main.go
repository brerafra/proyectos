package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

func makeRequest(client *http.Client, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("Request creation failed: %v", err)
	}
	req.Header.Set("Authorization", "Bearer your-token-here")

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 10 * time.Second

	return backoff.Retry(func() error {
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("Request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("Unexpected status: %v", resp.StatusCode)
		}
		return nil
	}, b)
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	err := makeRequest(client, "http://api.example.com/data")
	if err != nil {
		log.Fatalf("API call failed: %v", err)
	}
	log.Println("API Call successful! 🎉")
}
