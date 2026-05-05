package ports

import "context"

// PostSave is a port invoked by the posts application service around persistence.
// Implementations are composed in the posts module (e.g. derived fields + optional ES hook).
type PostSave interface {
	BeforePostSave(ctx context.Context, postID, slug string) error
	AfterPostSave(ctx context.Context, postID, slug string) error
}
