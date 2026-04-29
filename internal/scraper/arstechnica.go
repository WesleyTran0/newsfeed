package scraper

import (
	"errors"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WesleyTran0/newsfeed/pkg/models"
)

func ParseArsTechnicaSec(html string) []models.Article {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(errors.New("expected valid html; failed to parse given HTML to ParseArsTechnicaSec"))
	}

	var articles []models.Article

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		header := s.Find("h2 a").First()
		title := strings.TrimSpace(header.Text())
		url, exists := header.Attr("href")
		if !exists {
			panic("failed to find href in article header")
		}
		description := strings.TrimSpace(s.Find("p").Text())
		footer := s.Find("div span")
		author := strings.TrimSpace(footer.First().Text())
		date, exists := footer.Find("div span time").Attr("datetime")
		if !exists {
			panic("failed to find datetime in footer's time tag")
		}

		pubDate, err := time.Parse(time.RFC3339, date)
		if err != nil {
			panic("Failed to parse publish date")
		}

		articles = append(articles, models.Article{Title: title, Author: author, PublishedDate: pubDate, URL: url, Description: &description, FetchDate: time.Now()})
	})

	return articles
}
