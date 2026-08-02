package main

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func normalizeURL(inputURL string) (string, error) {
	if inputURL == "" {
		return "", errors.New("Invalid input")
	}
	url, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	noTrailingSlash := strings.TrimRight(url.Path, "/")
	url.Path = noTrailingSlash
	return url.Host + url.Path, nil
}

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set(req.UserAgent(), "BootCrawler/1.0")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 400 {
		return "", errors.New(resp.Status)
	}
	defer resp.Body.Close()
	return "", nil
}
