package scraper

import (
	"errors"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WesleyTran0/newsfeed/pkg/models"
)

func ParseTLDRSec(html string) []models.Article {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(errors.New("expected valid html; failed to parse given HTML to ParseArsTechnicaSec"))
	}

	var articles []models.Article

	doc.Find(`a[href^="/p/"]`).Not(":has(figure)").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Attr("href")
		if !exists {
			panic("Could not find the href of the \\<a\\> tag")
		}
		date, exists := s.Find("div span time").Attr("datetime")
		if !exists {
			panic("Could not find the datetime of the span time tags")
		}
		pubDate, err := time.Parse(time.RFC3339, date)
		if err != nil {
			panic(err)
		}
		title := strings.TrimSpace(s.Find("h2").Text())
		desc := strings.TrimSpace(s.Find("p").Text())
		author := strings.TrimSpace(s.Find(`a[href^="/authors"]`).Find("div span").Text())

		articles = append(articles, models.Article{Title: title, Author: author, PublishedDate: pubDate, URL: url, Description: &desc, FetchDate: time.Now()})
	})

	return articles
}
