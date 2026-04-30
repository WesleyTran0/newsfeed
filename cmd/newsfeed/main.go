package main

import (
	"flag"
	"log"
	"slices"
	"time"

	"github.com/WesleyTran0/newsfeed/internal/output"
	"github.com/WesleyTran0/newsfeed/internal/scraper"
	"github.com/WesleyTran0/newsfeed/pkg/models"
)

var newsSites = map[string]func(string) []models.Article{
	"https://arstechnica.com/security/":            scraper.ParseArsTechnicaSec,
	"https://this.weekinsecurity.com/articles/":    scraper.ParseThisWeekInSecurity,
	"https://this.weekinsecurity.com/past-issues/": scraper.ParseThisWeekInSecurity,
	"https://tldrsec.com/t/Summary":                scraper.ParseTLDRSec,
	"https://tldrsec.com/t/Blog":                   scraper.ParseTLDRSec,
}

var technologySites = map[string]func(string) []models.Article{
	"https://tldrsec.com/t/Newsletter": scraper.ParseTLDRSec,
}

func main() {
	path := flag.String("output", "./output.json", "The output file path to where this scraper will write to")
	flag.Parse()

	s := scraper.NewScraper(10 * time.Second)

	news := s.ScrapeAll(newsSites)
	technologies := s.ScrapeAll(technologySites)

	articles := append(news, technologies...)
	slices.SortFunc(articles, func(a1 models.Article, a2 models.Article) int {
		return a2.PublishedDate.Compare(a1.PublishedDate)
	})
	if err := output.WriteJSON(articles, *path); err != nil {
		log.Fatalf("failed to write output: %v", err)
	}
}
