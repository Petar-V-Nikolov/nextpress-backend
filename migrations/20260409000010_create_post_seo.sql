-- +migrate Up
CREATE TABLE IF NOT EXISTS post_seo (
    post_id         UUID PRIMARY KEY REFERENCES posts(id) ON DELETE CASCADE,
    title           TEXT NULL,
    description     TEXT NULL,
    canonical_url   TEXT NULL,
    robots          TEXT NULL,
    og_type         TEXT NULL,
    og_image_url    TEXT NULL,
    twitter_card    TEXT NULL,
    structured_data JSONB NOT NULL DEFAULT '{}'::jsonb,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS post_seo;

