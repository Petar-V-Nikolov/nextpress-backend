-- +migrate Up
CREATE TABLE IF NOT EXISTS post_metrics (
    post_id                   UUID PRIMARY KEY REFERENCES posts(id) ON DELETE CASCADE,
    word_count                INT NOT NULL DEFAULT 0,
    character_count           INT NOT NULL DEFAULT 0,
    reading_time_minutes      INT NOT NULL DEFAULT 0,
    est_read_time_seconds     INT NOT NULL DEFAULT 0,
    view_count                BIGINT NOT NULL DEFAULT 0,
    unique_visitors_7d        BIGINT NOT NULL DEFAULT 0,
    scroll_depth_avg_percent  REAL NOT NULL DEFAULT 0,
    bounce_rate_percent       REAL NOT NULL DEFAULT 0,
    avg_time_on_page_seconds  INT NOT NULL DEFAULT 0,
    comment_count             INT NOT NULL DEFAULT 0,
    like_count                INT NOT NULL DEFAULT 0,
    share_count               INT NOT NULL DEFAULT 0,
    bookmark_count            INT NOT NULL DEFAULT 0,
    updated_at                TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS post_metrics;

