package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/WesleyTran0/newsfeed/internal/scraper"
)

func main() {
	s := scraper.NewScraper(10 * time.Second)

	// articles, err := s.ScraperSource("https://arstechnica.com/security/", scraper.ParseArsTechnicaSec)
	// articles, err := s.ScraperSource("https://this.weekinsecurity.com/articles/", scraper.ParseThisWeekInSecurity)
	// if err != nil {
	// 	panic(err)
	// }
	// issues, err := s.ScraperSource("https://this.weekinsecurity.com/past-issues/", scraper.ParseThisWeekInSecurity)
	// if err != nil {
	// 	panic(err)
	// }
	newsletters, err := s.ScraperSource("https://tldrsec.com/t/Newsletter", scraper.ParseTLDRSec)
	if err != nil {
		panic(err)
	}
	summaries, err := s.ScraperSource("https://tldrsec.com/t/Summary", scraper.ParseTLDRSec)
	if err != nil {
		panic(err)
	}
	blogs, err := s.ScraperSource("https://tldrsec.com/t/Blog", scraper.ParseTLDRSec)
	if err != nil {
		panic(err)
	}

	all := slices.Concat(newsletters, summaries, blogs)

	fmt.Printf("%+v\n", all)
}
