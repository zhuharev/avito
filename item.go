package avito

import "time"

// Item represents an avito item
type Item struct {
	ID    int
	Title string
	URL   string

	Price int

	Published       time.Time
	PublishedString string

	Body string

	IsCompany bool
}
