package scraper

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/WesleyTran0/newsfeed/pkg/models"
)

func ParseArsTechnicaSec(html string) []models.Article {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		// TODO: todo
	}

	var articles []models.Article

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		header := s.Find("h2 a").First()
		title := strings.TrimSpace(header.Text())
		link, _ := header.Attr("href")
		description := strings.TrimSpace(s.Find("p").Text())
		footer := s.Find("div span")
		author := strings.TrimSpace(footer.First().Text())
		date, _ := footer.Find("div span time").Attr("datetime")

		fmt.Printf("title: %s\ndate: %s, author: %s\n", title, date, author)
		fmt.Printf("description: %s\nlink: %s\n\n", description, link)
	})

	return articles
}
