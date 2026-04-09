-- +migrate Up
CREATE TABLE IF NOT EXISTS series (
    id         UUID PRIMARY KEY,
    title      TEXT        NOT NULL,
    slug       TEXT        NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS post_series (
    post_id    UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    series_id  UUID NOT NULL REFERENCES series(id) ON DELETE CASCADE,
    part_index INT  NULL,
    part_label TEXT NULL,
    PRIMARY KEY (post_id, series_id)
);

CREATE INDEX IF NOT EXISTS idx_post_series_series_id ON post_series(series_id);

-- +migrate Down
DROP TABLE IF EXISTS post_series;
DROP TABLE IF EXISTS series;

