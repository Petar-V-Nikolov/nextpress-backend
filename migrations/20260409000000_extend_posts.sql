-- +migrate Up
ALTER TABLE posts
    ADD COLUMN IF NOT EXISTS uuid UUID NULL,
    ADD COLUMN IF NOT EXISTS post_type TEXT NULL,
    ADD COLUMN IF NOT EXISTS format TEXT NULL,
    ADD COLUMN IF NOT EXISTS subtitle TEXT NULL,
    ADD COLUMN IF NOT EXISTS excerpt TEXT NULL,
    ADD COLUMN IF NOT EXISTS visibility TEXT NOT NULL DEFAULT 'public',
    ADD COLUMN IF NOT EXISTS locale TEXT NOT NULL DEFAULT 'en-US',
    ADD COLUMN IF NOT EXISTS timezone TEXT NOT NULL DEFAULT 'UTC',
    ADD COLUMN IF NOT EXISTS scheduled_publish_at TIMESTAMPTZ NULL,
    ADD COLUMN IF NOT EXISTS first_indexed_at TIMESTAMPTZ NULL,
    ADD COLUMN IF NOT EXISTS reviewer_user_id UUID NULL REFERENCES users(id),
    ADD COLUMN IF NOT EXISTS last_edited_by_user_id UUID NULL REFERENCES users(id),
    ADD COLUMN IF NOT EXISTS workflow_stage TEXT NOT NULL DEFAULT 'draft',
    ADD COLUMN IF NOT EXISTS revision INT NOT NULL DEFAULT 1,
    ADD COLUMN IF NOT EXISTS custom_fields JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS flags JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS engagement JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS workflow JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS featured_media_id UUID NULL REFERENCES media(id),
    ADD COLUMN IF NOT EXISTS featured_alt TEXT NULL,
    ADD COLUMN IF NOT EXISTS featured_width INT NULL,
    ADD COLUMN IF NOT EXISTS featured_height INT NULL,
    ADD COLUMN IF NOT EXISTS featured_focal_x REAL NULL,
    ADD COLUMN IF NOT EXISTS featured_focal_y REAL NULL,
    ADD COLUMN IF NOT EXISTS featured_credit TEXT NULL,
    ADD COLUMN IF NOT EXISTS featured_license TEXT NULL;

CREATE UNIQUE INDEX IF NOT EXISTS idx_posts_uuid_unique ON posts(uuid) WHERE uuid IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_posts_visibility ON posts(visibility);
CREATE INDEX IF NOT EXISTS idx_posts_locale ON posts(locale);
CREATE INDEX IF NOT EXISTS idx_posts_scheduled_publish_at ON posts(scheduled_publish_at);

-- +migrate Down
DROP INDEX IF EXISTS idx_posts_scheduled_publish_at;
DROP INDEX IF EXISTS idx_posts_locale;
DROP INDEX IF EXISTS idx_posts_visibility;
DROP INDEX IF EXISTS idx_posts_uuid_unique;

ALTER TABLE posts
    DROP COLUMN IF EXISTS featured_license,
    DROP COLUMN IF EXISTS featured_credit,
    DROP COLUMN IF EXISTS featured_focal_y,
    DROP COLUMN IF EXISTS featured_focal_x,
    DROP COLUMN IF EXISTS featured_height,
    DROP COLUMN IF EXISTS featured_width,
    DROP COLUMN IF EXISTS featured_alt,
    DROP COLUMN IF EXISTS featured_media_id,
    DROP COLUMN IF EXISTS workflow,
    DROP COLUMN IF EXISTS engagement,
    DROP COLUMN IF EXISTS flags,
    DROP COLUMN IF EXISTS custom_fields,
    DROP COLUMN IF EXISTS revision,
    DROP COLUMN IF EXISTS workflow_stage,
    DROP COLUMN IF EXISTS last_edited_by_user_id,
    DROP COLUMN IF EXISTS reviewer_user_id,
    DROP COLUMN IF EXISTS first_indexed_at,
    DROP COLUMN IF EXISTS scheduled_publish_at,
    DROP COLUMN IF EXISTS timezone,
    DROP COLUMN IF EXISTS locale,
    DROP COLUMN IF EXISTS visibility,
    DROP COLUMN IF EXISTS excerpt,
    DROP COLUMN IF EXISTS subtitle,
    DROP COLUMN IF EXISTS format,
    DROP COLUMN IF EXISTS post_type,
    DROP COLUMN IF EXISTS uuid;

