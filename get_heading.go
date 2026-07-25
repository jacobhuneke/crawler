package main

import (
	"log"
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
