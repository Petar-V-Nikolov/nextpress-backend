package domain

import "context"

type Repository interface {
	Create(ctx context.Context, post *Post) error
	FindByID(ctx context.Context, id PostID) (*Post, error)
	FindBySlug(ctx context.Context, slug string) (*Post, error)
	List(ctx context.Context, includeDeleted bool, limit int, offset int) ([]Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id PostID) error
}

