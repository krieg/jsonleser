package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

// Anordnung represents configuration items
type Anordnung struct {
	BaseURL   string
	UploadURL string
}

func zurückholen(dateikennung string) (*Anordnung, error) {
	file, _ := os.Open(dateikennung)
	decoder := json.NewDecoder(file)
	anordnung := Anordnung{}
	err := decoder.Decode(&anordnung)
	if err != nil {
		return nil, err
	}
	return &anordnung, nil
}

func main() {
	baseURL, _ := url.Parse("defaultBase")
	uploadURL, _ := url.Parse("defaultUpload")
	a, err := zurückholen("domain.json")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		if a.BaseURL != "" {
			baseURL, _ = url.Parse(a.BaseURL)
		}
		if a.UploadURL != "" {
			uploadURL, _ = url.Parse(a.UploadURL)
		}
		fmt.Println(baseURL)
		fmt.Println(uploadURL)
	}
}