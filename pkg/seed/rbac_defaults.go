package seed

import (
	rbacp "github.com/nextpresskit/backend/internal/modules/rbac/persistence"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Deterministic UUIDs so environments stay consistent.
const (
	RoleAdminID = "00000000-0000-0000-0000-000000000001"

	PermissionAdminPingID       = "00000000-0000-0000-0000-000000000101"
	PermissionRBACManageID      = "00000000-0000-0000-0000-000000000102"
	PermissionPostsReadID       = "00000000-0000-0000-0000-000000000201"
	PermissionPostsWriteID      = "00000000-0000-0000-0000-000000000202"
	PermissionPagesReadID       = "00000000-0000-0000-0000-000000000203"
	PermissionPagesWriteID      = "00000000-0000-0000-0000-000000000204"
	PermissionCategoriesReadID  = "00000000-0000-0000-0000-000000000205"
	PermissionCategoriesWriteID = "00000000-0000-0000-0000-000000000206"
	PermissionTagsReadID        = "00000000-0000-0000-0000-000000000207"
	PermissionTagsWriteID       = "00000000-0000-0000-0000-000000000208"
	PermissionMediaReadID       = "00000000-0000-0000-0000-000000000209"
	PermissionMediaWriteID      = "00000000-0000-0000-0000-000000000210"
)

func knownPermissionID(code string) (string, bool) {
	m := map[string]string{
		"admin:ping":         PermissionAdminPingID,
		"rbac:manage":        PermissionRBACManageID,
		"posts:read":         PermissionPostsReadID,
		"posts:write":        PermissionPostsWriteID,
		"pages:read":         PermissionPagesReadID,
		"pages:write":        PermissionPagesWriteID,
		"categories:read":    PermissionCategoriesReadID,
		"categories:write":   PermissionCategoriesWriteID,
		"tags:read":          PermissionTagsReadID,
		"tags:write":         PermissionTagsWriteID,
		"media:read":         PermissionMediaReadID,
		"media:write":        PermissionMediaWriteID,
	}
	id, ok := m[code]
	return id, ok
}

// SeedRBACDefaults upserts the admin role, permission rows for known codes, and admin role links.
func SeedRBACDefaults(db *gorm.DB, permissionCodes []string) error {
	roles := []rbacp.Role{
		{ID: RoleAdminID, Name: "admin"},
	}
	for i := range roles {
		if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roles[i]).Error; err != nil {
			return err
		}
	}
	for _, code := range permissionCodes {
		id, ok := knownPermissionID(code)
		if !ok {
			continue
		}
		p := rbacp.Permission{ID: id, Code: code}
		if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&p).Error; err != nil {
			return err
		}
		link := rbacp.RolePermission{RoleID: RoleAdminID, PermissionID: id}
		if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&link).Error; err != nil {
			return err
		}
	}
	return nil
}
