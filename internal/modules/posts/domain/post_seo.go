package domain

import "context"

// PostSEOStore manages post_seo rows.
type PostSEOStore interface {
	DeleteSEO(ctx context.Context, postID PostID) error
	UpsertSEOOnly(ctx context.Context, postID PostID, seo *PostSEO) error
}
