package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %v\n", os.Args[1])
		base, err := url.Parse(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		conc, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		max, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}

		cfg := config{
			pages:              make(map[string]PageData),
			baseURL:            base,
			mu:                 &sync.Mutex{},
			concurrencyControl: make(chan struct{}, conc),
			wg:                 &sync.WaitGroup{},
			maxPages:           max,
		}

		cfg.wg.Add(1)
		go cfg.crawlPage(base.String())
		cfg.wg.Wait()
		writeJSONReport(cfg.pages, "report.json")
	}
}
