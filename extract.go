package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
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
