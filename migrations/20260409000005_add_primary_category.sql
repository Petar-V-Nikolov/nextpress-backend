-- +migrate Up
ALTER TABLE posts
    ADD COLUMN IF NOT EXISTS primary_category_id UUID NULL REFERENCES categories(id);

CREATE INDEX IF NOT EXISTS idx_posts_primary_category_id ON posts(primary_category_id);

-- +migrate Down
DROP INDEX IF EXISTS idx_posts_primary_category_id;
ALTER TABLE posts DROP COLUMN IF EXISTS primary_category_id;

