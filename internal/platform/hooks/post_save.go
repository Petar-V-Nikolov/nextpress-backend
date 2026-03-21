package hooks

import "context"

// PostSave is implemented by the plugin hook registry (or any adapter) and
// invoked by the posts application service around persistence. Keeping this
// interface in platform/hooks avoids a dependency from posts → plugins.
type PostSave interface {
	BeforePostSave(ctx context.Context, postID, slug string) error
	AfterPostSave(ctx context.Context, postID, slug string) error
}
