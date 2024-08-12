package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	url := os.Args[1]
	isUsingGoogleAnalytics, err := isUsingGoogleAnalytics(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Is the website using Google Analytics? %v\n", isUsingGoogleAnalytics)
}

func isUsingGoogleAnalytics(url string) (bool, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(resp.Body)

	htmlContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// List of patterns that are used by Google Analytics, add more patterns if needed
	googleAnalyticsPatterns := []string{
		"google-analytics.com/ga.js",
		"google-analytics.com/analytics.js",
		"google-analytics.com/gtag/js",
		"google-analytics.com/urchin.js",
		"google-analytics.com/ga.js",
		"google-analytics.com/dc.js",
		"googletagmanager.com/gtag/js",
		"gtag('config'",
	}

	for _, pattern := range googleAnalyticsPatterns {
		if strings.Contains(string(htmlContent), pattern) {
			return true, nil
		}
	}

	return false, nil
}
