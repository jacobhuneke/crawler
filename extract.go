package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func extractPageData(html, pageURL string) PageData {
	var data PageData
	url, err := url.Parse(pageURL)
	if err != nil {
		log.Fatal(err)
	}

	data.URL = pageURL
	data.Heading = getHeadingFromHTML(html)
	data.FirstParagraph = getFirstParagraphFromHTML(html)
	data.OutgoingLinks, err = getURLsFromHTML(html, url)
	if err != nil {
		log.Fatal(err)
	}
	data.ImageURLs, err = getImagesFromHTML(html, url)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
