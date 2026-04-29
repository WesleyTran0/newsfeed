package main

import (
	"fmt"
	"time"

	"github.com/WesleyTran0/newsfeed/internal/scraper"
)

func main() {
	s := scraper.NewScraper(10 * time.Second)
	html, _ := s.Fetch("https://arstechnica.com/security/")

	articles := scraper.ParseArsTechnicaSec(html)
	fmt.Printf("%s", articles)
}
