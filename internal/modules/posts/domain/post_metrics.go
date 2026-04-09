package domain

import "context"

// PostMetricsStore reads/writes post_metrics (read path; writes often via full post upsert).
type PostMetricsStore interface {
	GetMetrics(ctx context.Context, postID PostID) (*PostMetrics, error)
}
