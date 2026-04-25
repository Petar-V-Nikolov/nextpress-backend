-- +migrate Up
DELETE FROM role_permissions
WHERE permission_id IN (
    SELECT id
    FROM permissions
    WHERE code IN ('menus:read', 'menus:write')
);

DELETE FROM permissions
WHERE code IN ('menus:read', 'menus:write');

DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS menus;

-- +migrate Down
CREATE TABLE IF NOT EXISTS menus (
    id         UUID PRIMARY KEY,
    name       TEXT        NOT NULL,
    slug       TEXT        NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS menu_items (
    id          UUID PRIMARY KEY,
    menu_id     UUID        NOT NULL REFERENCES menus(id) ON DELETE CASCADE,
    parent_id   UUID        NULL REFERENCES menu_items(id) ON DELETE CASCADE,
    label       TEXT        NOT NULL,
    item_type   TEXT        NOT NULL,
    ref_id      UUID        NULL,
    url         TEXT        NULL,
    sort_order  INT         NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_menu_items_menu_id ON menu_items(menu_id);
CREATE INDEX IF NOT EXISTS idx_menu_items_parent_id ON menu_items(parent_id);
CREATE INDEX IF NOT EXISTS idx_menu_items_sort_order ON menu_items(menu_id, sort_order);
