package models

import (
	"time"
)

type Article struct {
	Title         string
	Author        string
	PublishedDate time.Time
	URL           string
	Description   *string
	FetchDate     time.Time
}
