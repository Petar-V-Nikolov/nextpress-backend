package domain

import "context"

// PostCoauthorsStore manages post_coauthors.
type PostCoauthorsStore interface {
	ReplaceCoauthors(ctx context.Context, postID PostID, userIDs []string) error
}
