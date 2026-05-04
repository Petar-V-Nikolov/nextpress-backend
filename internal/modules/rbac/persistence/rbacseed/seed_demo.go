package rbacseed

import (
	"fmt"
	"time"

	rbacp "github.com/nextpresskit/backend/internal/modules/rbac/persistence"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	demoSeedRows               = 100
	demoSeedDefaultPermissions = 13
	superadminRoleName         = "superadmin"
	superadminUserUUID         = "00000000-0000-0000-0100-000000000001"
	superadminRoleUUID         = "00000000-0000-0000-0200-000000000002"
	roleAdminSeedID            = "00000000-0000-0000-0000-000000000001"
)

// SeedDemo inserts demo roles, synthetic permissions, synthetic role-permission links, and user_roles for the superadmin.
func SeedDemo(tx *gorm.DB) error {
	if err := seedRoles(tx); err != nil {
		return err
	}
	if err := seedExtraPermissions(tx); err != nil {
		return err
	}
	if err := seedSyntheticRolePermissions(tx); err != nil {
		return err
	}
	return seedUserRoles(tx)
}

func seedRoles(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		id := seedUUID(0x0200, i)
		name := fmt.Sprintf("role-%03d", i)
		if i == 1 {
			id = roleAdminSeedID
			name = "admin"
		} else if i == 2 {
			id = superadminRoleUUID
			name = superadminRoleName
		}
		r := rbacp.Role{ID: id, Name: name}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "name"}},
			DoUpdates: clause.Assignments(map[string]any{
				"updated_at": time.Now().UTC(),
			}),
		}).Create(&r).Error; err != nil {
			return fmt.Errorf("roles row %d: %w", i, err)
		}
	}
	return nil
}

func seedExtraPermissions(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows-demoSeedDefaultPermissions; i++ {
		p := rbacp.Permission{
			ID:   seedUUID(0x0300, i),
			Code: fmt.Sprintf("seed:permission:%03d", i),
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "code"}},
			DoUpdates: clause.Assignments(map[string]any{
				"updated_at": time.Now().UTC(),
			}),
		}).Create(&p).Error; err != nil {
			return fmt.Errorf("permissions row %d: %w", i, err)
		}
	}
	return nil
}

func seedSyntheticRolePermissions(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows-demoSeedDefaultPermissions; i++ {
		var permID, roleID string
		code := fmt.Sprintf("seed:permission:%03d", i)
		roleName := fmt.Sprintf("role-%03d", i+demoSeedDefaultPermissions)
		if err := tx.Model(&rbacp.Permission{}).Select("id").Where("code = ?", code).Scan(&permID).Error; err != nil {
			return err
		}
		if err := tx.Model(&rbacp.Role{}).Select("id").Where("name = ?", roleName).Scan(&roleID).Error; err != nil {
			return err
		}
		if permID == "" || roleID == "" {
			continue
		}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&rbacp.RolePermission{
			RoleID: roleID, PermissionID: permID,
		}).Error; err != nil {
			return fmt.Errorf("role_permissions row %d: %w", i, err)
		}
	}
	return nil
}

func seedUserRoles(tx *gorm.DB) error {
	superPub := userPublicIDFromUUID(tx, "users", superadminUserUUID)
	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&rbacp.UserRole{
		UserID: superPub,
		RoleID: superadminRoleUUID,
	}).Error; err != nil {
		return fmt.Errorf("user_roles superadmin link: %w", err)
	}
	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&rbacp.UserRole{
		UserID: superPub,
		RoleID: roleAdminSeedID,
	}).Error; err != nil {
		return fmt.Errorf("user_roles admin link: %w", err)
	}
	for i := 2; i <= 99; i++ {
		ur := rbacp.UserRole{
			UserID: userPublicIDFromUUID(tx, "users", seedUUID(0x0100, i)),
			RoleID: seedUUID(0x0200, i+1),
		}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&ur).Error; err != nil {
			return fmt.Errorf("user_roles row %d: %w", i, err)
		}
	}
	return nil
}

func seedUUID(namespace, index int) string {
	return fmt.Sprintf("00000000-0000-0000-%04x-%012x", namespace, index)
}

func userPublicIDFromUUID(tx *gorm.DB, table, userUUID string) int64 {
	var id int64
	_ = tx.Table(table).Select("public_id").Where("id = ?", userUUID).Scan(&id).Error
	return id
}
