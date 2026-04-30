package models

import (
	"time"
)

type Article struct {
	Title         string    `json:"title"`
	Author        string    `json:"author(s)"`
	PublishedDate time.Time `json:"date"`
	URL           string    `json:"url"`
	Description   *string   `json:"description"`
	FetchDate     time.Time `json:"dateFetched"`
}
