package domain

import "context"

// PostSeriesLinkStore manages post_series membership.
type PostSeriesLinkStore interface {
	SetPostSeries(ctx context.Context, postID PostID, seriesID *string, partIndex *int, partLabel *string) error
}
