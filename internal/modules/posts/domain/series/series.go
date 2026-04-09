package series

import (
	"context"
	"time"
)

// Series is a curated post series entity.
type Series struct {
	ID        string
	Title     string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// SeriesRepository is top-level CRUD for series.
type SeriesRepository interface {
	ListSeries(ctx context.Context) ([]Series, error)
	CreateSeries(ctx context.Context, s *Series) error
	FindSeriesByID(ctx context.Context, id string) (*Series, error)
	UpdateSeries(ctx context.Context, s *Series) error
	DeleteSeries(ctx context.Context, id string) error
}
