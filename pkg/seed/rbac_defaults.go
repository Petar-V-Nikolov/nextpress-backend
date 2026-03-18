package seed

import "gorm.io/gorm"

// Deterministic UUIDs so environments stay consistent.
const (
	RoleAdminID = "00000000-0000-0000-0000-000000000001"

	PermissionAdminPingID = "00000000-0000-0000-0000-000000000101"
	PermissionRBACManageID = "00000000-0000-0000-0000-000000000102"
)

func SeedRBACDefaults(db *gorm.DB) error {
	// roles
	if err := db.Exec(
		`INSERT INTO roles (id, name) VALUES (?, ?) ON CONFLICT (name) DO NOTHING`,
		RoleAdminID, "admin",
	).Error; err != nil {
		return err
	}

	// permissions
	if err := db.Exec(
		`INSERT INTO permissions (id, code) VALUES (?, ?) ON CONFLICT (code) DO NOTHING`,
		PermissionAdminPingID, "admin:ping",
	).Error; err != nil {
		return err
	}
	if err := db.Exec(
		`INSERT INTO permissions (id, code) VALUES (?, ?) ON CONFLICT (code) DO NOTHING`,
		PermissionRBACManageID, "rbac:manage",
	).Error; err != nil {
		return err
	}

	// role_permissions links
	if err := db.Exec(
		`INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?) ON CONFLICT DO NOTHING`,
		RoleAdminID, PermissionAdminPingID,
	).Error; err != nil {
		return err
	}
	if err := db.Exec(
		`INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?) ON CONFLICT DO NOTHING`,
		RoleAdminID, PermissionRBACManageID,
	).Error; err != nil {
		return err
	}

	return nil
}

