package domain

import "context"

// PostLoadUpdater is the minimal persistence surface for PostSave hooks that reload and patch a post.
type PostLoadUpdater interface {
	FindByID(ctx context.Context, id PostID) (*Post, error)
	Update(ctx context.Context, post *Post) error
}
