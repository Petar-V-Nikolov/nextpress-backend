package domain

import "context"

// PostWriter persists core post rows.
type PostWriter interface {
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id PostID) error
}
