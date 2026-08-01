package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	header := doc.Find("h1").Text()
	if header == "" {
		header = doc.Find("h2").Text()
	}

	return header
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	main := doc.Find("main")
	paragraph := main.Find("p")
	first := paragraph.First().Text()
	if first != "" {
		return first
	} else {
		p := doc.Find("p").Text()
		return p
	}

}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	urls := []string{}
	var urlErr error
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if exists == true {
			newURL, err := url.Parse(val)
			if err != nil {
				urlErr = err
			}
			urls = append(urls, baseURL.ResolveReference(newURL).String())
		}
	})
	if urlErr != nil {
		return nil, urlErr
	}
	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}
	imgs := []string{}
	var imgErr error
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		val, exists := s.Attr("src")
		if exists == true {
			newURL, err := url.Parse(val)
			if err != nil {
				imgErr = err
			}
			imgs = append(imgs, baseURL.ResolveReference(newURL).String())
		}
	})
	if imgErr != nil {
		return nil, imgErr
	}
	return imgs, nil
}
