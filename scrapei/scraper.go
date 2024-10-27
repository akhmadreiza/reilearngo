package scrapei

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// News holds the scraped data
type News struct {
	URL   string
	Title string
}

func FetchUrl(url string) ([]News, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []News{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []News{}, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []News{}, err
	}

	var listNews []News
	doc.Find("article.list-content__item").Each(func(i int, s *goquery.Selection) {
		news := News{}
		media := s.Find("h3.media__title")

		var resultUrl string
		newsUrl, attrExists := media.Find("a").Attr("href")
		if attrExists {
			resultUrl = newsUrl
		}

		newsTitle := media.Find("a").Text()

		news.URL = resultUrl
		news.Title = newsTitle

		listNews = append(listNews, news)
	})

	return listNews, nil
}
