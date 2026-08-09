package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.baseURL.Host != currentURL.Host {
		return
	}

	curr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	if !cfg.addPageVisit(curr) {
		return
	} else {
		html, err := getHTML(rawCurrentURL)
		if err != nil {
			log.Fatal(err)
		}

		cfg.mu.Lock()
		cfg.pages[curr] = extractPageData(html, curr)
		cfg.mu.Unlock()

		urls, err := getURLsFromHTML(html, cfg.baseURL)
		if err != nil {
			log.Fatal(err)
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
