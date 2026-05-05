// Package seed runs database seeders (reference data and optional dev data).
package seed

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/nextpresskit/backend/internal/appregistry"
	"github.com/nextpresskit/backend/internal/kit"
)

// Run runs RBAC defaults then module seeds in registry order inside one transaction.
func Run(db *gorm.DB) error {
	zl, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("logger: %w", err)
	}
	logger := zl.Sugar()
	defer func() { _ = zl.Sync() }()

	modules := kit.ResolveModulesFromRegistry(logger, appregistry.ModuleRegistry())
	codes := kit.CollectPermissionCodes(modules)
	if err := SeedRBACDefaults(db, codes); err != nil {
		return fmt.Errorf("seed rbac_defaults: %w", err)
	}
	opts := kit.SeedOpts{DefaultPermissionCodes: codes}
	return db.Transaction(func(tx *gorm.DB) error {
		for _, m := range modules {
			log.Printf("seeding module %s...", m.ID())
			if err := m.Seed(tx, opts); err != nil {
				return fmt.Errorf("seed %s: %w", m.ID(), err)
			}
			log.Printf("seeded module %s ok", m.ID())
		}
		return nil
	})
}
