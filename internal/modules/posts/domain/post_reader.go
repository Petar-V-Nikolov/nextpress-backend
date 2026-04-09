package domain

import "context"

// PostReader loads posts (admin and public list/read paths).
type PostReader interface {
	FindByID(ctx context.Context, id PostID) (*Post, error)
	FindBySlug(ctx context.Context, slug string) (*Post, error)
	List(ctx context.Context, includeDeleted bool, limit int, offset int) ([]Post, error)
	ListFiltered(ctx context.Context, includeDeleted bool, limit int, offset int, status string, authorID string, q string) ([]Post, error)
	ListPublished(ctx context.Context, limit int, offset int, q string, categoryID string, tagID string) ([]Post, error)
	FindPublishedBySlug(ctx context.Context, slug string) (*Post, error)
}
