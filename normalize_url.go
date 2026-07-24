package main

import (
	"errors"
	"net/url"
	"strings"
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
