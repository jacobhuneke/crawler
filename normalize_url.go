package main

import (
	"errors"
	"io"
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

	req.Header.Set("User-Agent", "BootCrawler/1.0")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return "", errors.New(resp.Status)
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("invalid content type")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
