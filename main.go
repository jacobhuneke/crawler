package main

import (
	"fmt"
	"os"
)

func main() {
	pages := make(map[string]int)
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %v\n", os.Args[1:])
		fmt.Println(getHTML(os.Args[1]))
		crawlPage(os.Args[1], os.Args[1], pages)
		for key, val := range pages {
			fmt.Printf("%s: %v\n", key, val)
		}
	}
}
