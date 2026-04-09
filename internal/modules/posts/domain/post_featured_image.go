package domain

import "context"

// PostFeaturedImageStore updates featured media columns on posts.
type PostFeaturedImageStore interface {
	SetFeaturedImage(ctx context.Context, postID PostID, mediaID *string, alt *string, width *int, height *int, focalX *float32, focalY *float32, credit *string, license *string) error
}
