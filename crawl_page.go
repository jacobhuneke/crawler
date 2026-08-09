package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	curr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !cfg.addPageVisit(curr) {
		return
	} else {
		fmt.Printf("crawling page %s\n", rawCurrentURL)
		html, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Println(err)
			return
		}

		cfg.mu.Lock()
		cfg.pages[curr] = extractPageData(html, curr)
		cfg.mu.Unlock()

		urls, err := getURLsFromHTML(html, cfg.baseURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, url := range urls {
			cfg.wg.Add(1)
			go cfg.crawlPage(url)
		}
	}

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, ok := cfg.pages[normalizedURL]
	if !ok {
		cfg.pages[normalizedURL] = PageData{}
		return true
	}

	return false
}
