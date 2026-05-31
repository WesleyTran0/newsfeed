package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/WesleyTran0/newsfeed/pkg/models"
)

type Scraper struct {
	client *http.Client
}

type FetchError struct {
	URL  string
	Err  error
	Desc *string
}

func (e *FetchError) Error() string {
	base := "could not fetch " + e.URL + ": " + e.Err.Error()
	if e.Desc != nil {
		return base + "\n" + *e.Desc
	}
	return base
}

func NewScraper(timeout time.Duration) *Scraper {
	return &Scraper{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Fetch scrapes the content from the given `url`
func (s *Scraper) Fetch(url string) (string, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return "", &FetchError{url, err, nil}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errMsg := fmt.Sprint("Expected StatusCode 200, Found: %i", resp.StatusCode)
		return "", &FetchError{url, err, &errMsg}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := "Failed to parse the body into a string"
		return "", &FetchError{url, err, &errMsg}
	}

	return string(body), nil
}

// ScraperSource returns the list of articles from the url. The url is parsed using the parser function.
// This errors if at any point, the parser fails
func (s *Scraper) ScraperSource(url string, parser func(string) []models.Article) ([]models.Article, error) {
	html, err := s.Fetch(url)
	if err != nil {
		return nil, err
	}

	return parser(html), nil
}

// ScrapeAll returns the list of all articles scraped from sources. Each source is parsed using its associated function.
func (s *Scraper) ScrapeAll(sources map[string]func(string) []models.Article) []models.Article {
	results := make(chan []models.Article, len(sources))

	var wg sync.WaitGroup

	for url, parser := range sources {
		wg.Go(func() {
			articles, err := s.ScraperSource(url, parser)
			if err != nil {
				log.Printf("error scraping %s: %v", url, err)
				results <- []models.Article{}
				return
			}
			results <- articles
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var all []models.Article
	for articles := range results {
		all = append(all, articles...)
	}
	return deduplicate(all)
}

// deduplicate removes all duplicate articles from articles. Duplicates are determined by the same url.
func deduplicate(articles []models.Article) []models.Article {
	seen := map[string]bool{}
	result := []models.Article{}

	for _, a := range articles {
		if !seen[a.URL] {
			seen[a.URL] = true
			result = append(result, a)
		}
	}
	return result
}
