package domain

import "context"

// SeriesRepository is top-level CRUD for series.
type SeriesRepository interface {
	ListSeries(ctx context.Context) ([]Series, error)
	CreateSeries(ctx context.Context, s *Series) error
	FindSeriesByID(ctx context.Context, id string) (*Series, error)
	UpdateSeries(ctx context.Context, s *Series) error
	DeleteSeries(ctx context.Context, id string) error
}
