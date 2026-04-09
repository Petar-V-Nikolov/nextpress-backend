package domain

import "context"

// PostGalleryStore manages post_gallery_items.
type PostGalleryStore interface {
	CreateGalleryItem(ctx context.Context, postID PostID, mediaID string, sortOrder int, caption *string, alt *string) (itemID string, err error)
	UpdateGalleryItem(ctx context.Context, postID PostID, itemID string, sortOrder *int, caption *string, alt *string) error
	DeleteGalleryItem(ctx context.Context, postID PostID, itemID string) error
}
