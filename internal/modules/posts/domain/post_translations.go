package domain

import "context"

// PostTranslationsStore manages post_translations and translation group linkage.
type PostTranslationsStore interface {
	PutPostTranslation(ctx context.Context, postID PostID, groupID *string, locale string) (resolvedGroupID string, err error)
	ClearPostTranslation(ctx context.Context, postID PostID) error
}
