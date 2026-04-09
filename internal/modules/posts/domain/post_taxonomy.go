package domain

import "context"

// PostTaxonomyWriter assigns taxonomy and primary category on posts.
type PostTaxonomyWriter interface {
	SetCategories(ctx context.Context, postID PostID, categoryIDs []string) error
	SetTags(ctx context.Context, postID PostID, tagIDs []string) error
	SetPrimaryCategory(ctx context.Context, postID PostID, categoryID *string) error
}
