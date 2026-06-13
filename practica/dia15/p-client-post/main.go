package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

func uploadFile(client *http.Client, url, content string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "example.txt")
	if err != nil {
		return fmt.Errorf("Form file creation failed: %v", err)
	}
	io.WriteString(part, content)
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("Request creation failed: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("REquest Failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status: %v", resp.Status)
	}
	return nil
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	err := uploadFile(client, "https://api.example.com/upload", "Hello, world!")
	if err != nil {
		log.Fatalf("Upload failed: %v", err)
	}
	log.Println("File uploaded successfully!")
}
