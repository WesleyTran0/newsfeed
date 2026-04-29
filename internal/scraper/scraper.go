package scraper

import (
	"fmt"
	"io"
	"net/http"
	"time"
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
		errMsg := fmt.Sprint("Expected StatusCode 200, Found: %d", resp.StatusCode)
		return "", &FetchError{url, err, &errMsg}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := "Failed to parse the body into a string"
		return "", &FetchError{url, err, &errMsg}
	}

	return string(body), nil
}
