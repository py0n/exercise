package tour

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type fetcher interface {
	fetch(url string) (body string, urls []string, err error)
}

type defaultFetcher struct{}

func (f defaultFetcher) fetch(url string) (string, []string, error) {
	var urls []string

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", nil, err
	}

	body, err := doc.Html()
	if err != nil {
		return "", nil, err
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, exists := s.Attr("href")
		if exists {
			urls = append(urls, url)
		}
	})

	return body, urls, nil
}

func crawl(url string, depth int, fetcher fetcher) error {
	if depth <= 0 {
		return nil
	}
	_, urls, err := fetcher.fetch(url)
	if err != nil {
		return err
	}
	fmt.Printf("found: %s %#v\n", url, urls)
	for _, u := range urls {
		crawl(u, depth-1, fetcher)
	}
	return nil
}

// C71 ...
func C71(opts Options, args []string) int {
	if len(args) != 3 {
		return 1
	}
	n, err := strconv.Atoi(args[2])
	if err != nil {
		return 1
	}
	var fetcher defaultFetcher
	crawl(args[1], n, fetcher)
	return 0
}
