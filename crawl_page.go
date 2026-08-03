package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}
	if baseURL.Host != currentURL.Host {
		return
	}
	curr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	_, exists := pages[curr]
	if exists {
		pages[curr] += 1
		return
	} else {
		pages[curr] = 1
		html, err := getHTML(rawCurrentURL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(html)
		urls, err := getURLsFromHTML(html, baseURL)
		if err != nil {
			log.Fatal(err)
		}
		for _, url := range urls {
			crawlPage(rawBaseURL, url, pages)
		}
	}

}
