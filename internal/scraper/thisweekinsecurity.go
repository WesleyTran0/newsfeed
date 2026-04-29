package scraper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/WesleyTran0/newsfeed/pkg/models"
)

func ParseThisWeekInSecurity(html string) []models.Article {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(errors.New("expected valid html; failed to parse given HTML to ParseArsTechnicaSec"))
	}

	var articles []models.Article

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h2.feed-title").Text())
		day := strings.TrimSpace(s.Find("div.feed-calendar-day").Text())
		month := strings.TrimSpace(s.Find("div.feed-calendar-month").Text())
		year := "20" + strings.TrimSpace(s.Find("div.feed-calendar-year").Text())[1:]
		date, err := time.Parse("2006-Jan-02", fmt.Sprintf("%s-%s-%s", year, month, day))
		if err != nil {
			fmt.Println("parse error:", err)
		}
		desc := strings.TrimSpace(s.Find("div.feed-excerpt").Text())
		url, exists := s.Find("a").Attr("href")
		if !exists {
			panic("failed to find href in article's a tag")
		}

		articles = append(articles, models.Article{Title: title, Author: "Zack Whittaker", PublishedDate: date, URL: "https://this.weekinsecurity.com" + url, Description: &desc, FetchDate: time.Now()})
	})
	return articles
}
