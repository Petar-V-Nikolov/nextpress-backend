package domain

import "context"

// PostSyndicationStore manages post_syndication rows.
type PostSyndicationStore interface {
	CreateSyndication(ctx context.Context, postID PostID, platform, url, status string) (id string, err error)
	UpdateSyndication(ctx context.Context, postID PostID, id string, platform, url, status *string) error
	DeleteSyndication(ctx context.Context, postID PostID, id string) error
	UpdateSyndicationByID(ctx context.Context, id string, platform, url, status *string) error
	DeleteSyndicationByID(ctx context.Context, id string) error
}
