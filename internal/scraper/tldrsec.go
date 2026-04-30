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

	doc.Find("div.w-full.p-3 div.space-y-3").Each(func(i int, s *goquery.Selection) {
		url, exists := s.Find(`a[href^="/p/"]`).Attr("href")
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
		// FIX: WRONG -> cannot look for /p/ since this authors <a> is not part of that. This currently makes the author of the most recent article, the author for the entire page
		author := strings.TrimSpace(s.Find(`a[href^="/authors"]`).First().Find("div span").Text())

		articles = append(articles, models.Article{Title: title, Author: author, PublishedDate: pubDate, URL: "https://tldrsec.com" + url, Description: &desc, FetchDate: time.Now()})
	})

	return articles
}
