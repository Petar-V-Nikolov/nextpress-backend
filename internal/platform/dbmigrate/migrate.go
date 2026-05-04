package dbmigrate

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/nextpresskit/backend/internal/kit"
)

// AutoMigrateAll runs module AutoMigrate hooks in invocation order (caller must pass FK-safe order).
func AutoMigrateAll(db *gorm.DB, modules []kit.Module) error {
	for _, m := range modules {
		if err := m.AutoMigrate(db); err != nil {
			return fmt.Errorf("%s: %w", m.ID(), err)
		}
	}
	ensureUserPublicIDDefault(db)
	return nil
}

func ensureUserPublicIDDefault(db *gorm.DB) {
	if err := db.Exec(`CREATE SEQUENCE IF NOT EXISTS users_public_id_seq`).Error; err != nil {
		return
	}
	_ = db.Exec(`SELECT setval('users_public_id_seq', GREATEST(COALESCE((SELECT MAX(public_id) FROM users), 0), 1))`).Error
	_ = db.Exec(`ALTER TABLE users ALTER COLUMN public_id SET DEFAULT nextval('users_public_id_seq'::regclass)`).Error
}

// DropPublicSchema drops every table in public (dev reset). Destructive.
func DropPublicSchema(db *gorm.DB) error {
	return db.Exec(`
		DO $$ DECLARE r RECORD;
		BEGIN
		  FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public')
		  LOOP
		    EXECUTE 'DROP TABLE IF EXISTS public.' || quote_ident(r.tablename) || ' CASCADE';
		  END LOOP;
		END $$;
	`).Error
}
