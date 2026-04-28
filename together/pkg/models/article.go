package models

import (
	"time"
)

type Article struct {
	Title         string
	PublishedDate time.Time
	URL           string
	FetchDate     time.Time
}
