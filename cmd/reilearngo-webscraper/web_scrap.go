package main

import (
	"fmt"

	"github.com/akhmadreiza/reilearngo/scrapei"
)

func main() {
	url := "https://www.detik.com/search/searchall?query=prabowo subianto&page=1&result_type=latest"
	//body, err := scrapei.FetchUrl(url)
	news, err := scrapei.FetchUrl(url)
	if err != nil {
		fmt.Println("Error fetching url:", err)
		return
	}

	for _, eachNews := range news {
		fmt.Printf("Scraped News: URL=%s, Title=%s\n", eachNews.URL, eachNews.Title)
	}
}
