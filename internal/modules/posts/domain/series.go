package domain

import "time"

type Series struct {
	ID        string
	Title     string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
