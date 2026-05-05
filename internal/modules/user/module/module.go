package module

import (
	"context"

	"github.com/nextpresskit/backend/internal/kit"
	userInfra "github.com/nextpresskit/backend/internal/modules/user/infrastructure"
	userp "github.com/nextpresskit/backend/internal/modules/user/persistence"
	"gorm.io/gorm"
)

type userMod struct{}

func (userMod) ID() string { return "user" }

func (userMod) Prepare(d *kit.Deps) error {
	d.UserRepo = userInfra.NewGormRepository(d.DB)
	return nil
}

func (userMod) RegisterAuth(*kit.Deps) error   { return nil }
func (userMod) RegisterPublic(*kit.Deps) error { return nil }
func (userMod) RegisterAdmin(*kit.Deps) error { return nil }

func (userMod) AutoMigrate(db *gorm.DB) error {
	return userp.AutoMigrate(db)
}

func (userMod) Seed(db *gorm.DB, _ kit.SeedOpts) error {
	return userp.SeedDemo(db)
}

func (userMod) Start(context.Context, *kit.Deps) error { return nil }

func (userMod) Permissions() []string { return nil }

// Module is the user slice (persistence + demo seed).
var Module kit.Module = userMod{}
