package domain

import "context"

// PostChangelogStore manages post_changelog.
type PostChangelogStore interface {
	CreateChangelog(ctx context.Context, postID PostID, userID *string, note string) (changelogID string, err error)
	DeleteChangelog(ctx context.Context, postID PostID, changelogID string) error
}
