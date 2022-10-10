package main

import (
	"GoogleScraper/internal/ScrapeAndParse"
	"fmt"
)

func main() {
	res, err := ScrapeAndParse.GoogleScrape("golang project", "com", "en", nil, 1, 30, 10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}

}
