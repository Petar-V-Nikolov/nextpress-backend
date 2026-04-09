-- +migrate Up
CREATE TABLE IF NOT EXISTS post_coauthors (
    post_id    UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    sort_order INT  NOT NULL DEFAULT 0,
    PRIMARY KEY (post_id, user_id)
);

CREATE TABLE IF NOT EXISTS post_gallery_items (
    id         UUID PRIMARY KEY,
    post_id    UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    media_id   UUID NOT NULL REFERENCES media(id),
    sort_order INT  NOT NULL DEFAULT 0,
    caption    TEXT NULL,
    alt        TEXT NULL
);
CREATE INDEX IF NOT EXISTS idx_post_gallery_items_post_id ON post_gallery_items(post_id);

CREATE TABLE IF NOT EXISTS post_changelog (
    id      UUID PRIMARY KEY,
    post_id UUID        NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_id UUID        NULL REFERENCES users(id),
    note    TEXT        NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_post_changelog_post_id ON post_changelog(post_id);

CREATE TABLE IF NOT EXISTS post_syndication (
    id         UUID PRIMARY KEY,
    post_id    UUID        NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    platform   TEXT        NOT NULL,
    url        TEXT        NOT NULL,
    status     TEXT        NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_post_syndication_post_id ON post_syndication(post_id);

CREATE TABLE IF NOT EXISTS translation_groups (
    id         UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS post_translations (
    post_id  UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    group_id UUID NOT NULL REFERENCES translation_groups(id) ON DELETE CASCADE,
    locale   TEXT NOT NULL,
    PRIMARY KEY (post_id)
);
CREATE INDEX IF NOT EXISTS idx_post_translations_group_id ON post_translations(group_id);

-- +migrate Down
DROP TABLE IF EXISTS post_translations;
DROP TABLE IF EXISTS translation_groups;
DROP TABLE IF EXISTS post_syndication;
DROP TABLE IF EXISTS post_changelog;
DROP TABLE IF EXISTS post_gallery_items;
DROP TABLE IF EXISTS post_coauthors;

