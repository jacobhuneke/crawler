package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	cfg := config{
		pages:              make(map[string]PageData),
		baseURL:            nil,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync.WaitGroup{},
	}

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {

		fmt.Printf("starting crawl of: %v\n", os.Args[1:])
		base, err := url.Parse(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		cfg.baseURL = base

		cfg.wg.Add(1)
		go cfg.crawlPage(base.String())
		cfg.wg.Wait()
	}
}
