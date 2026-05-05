package module

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nextpresskit/backend/internal/kit"
	rbacapp "github.com/nextpresskit/backend/internal/modules/rbac/application"
	rbacinfra "github.com/nextpresskit/backend/internal/modules/rbac/infrastructure"
	rbacp "github.com/nextpresskit/backend/internal/modules/rbac/persistence"
	rbacseed "github.com/nextpresskit/backend/internal/modules/rbac/persistence/rbacseed"
	rbactransport "github.com/nextpresskit/backend/internal/modules/rbac/transport"
	platformMiddleware "github.com/nextpresskit/backend/internal/platform/middleware"
	"gorm.io/gorm"
)

type rbacMod struct{}

func (rbacMod) ID() string { return "rbac" }

func (rbacMod) Prepare(d *kit.Deps) error {
	d.RBACRepo = rbacinfra.NewGormRepository(d.DB)
	d.PermissionChecker = rbacinfra.NewGormPermissionChecker(d.DB)
	d.RBACService = rbacapp.NewService(d.RBACRepo)
	d.RBACHandler = rbactransport.NewHandler(d.RBACService)
	return nil
}

func (rbacMod) RegisterAuth(*kit.Deps) error { return nil }

func (rbacMod) RegisterPublic(*kit.Deps) error { return nil }

func (rbacMod) RegisterAdmin(d *kit.Deps) error {
	adminManagement := d.Admin.Group("")
	adminManagement.Use(platformMiddleware.RequirePermission(d.PermissionChecker, "rbac:manage"))
	d.RBACHandler.RegisterRoutes(adminManagement)

	d.Admin.GET("/ping",
		platformMiddleware.RequirePermission(d.PermissionChecker, "admin:ping"),
		func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": true})
		},
	)

	if d.RBACCfg.BootstrapEnabled {
		d.Admin.POST("/bootstrap/claim-admin", func(c *gin.Context) {
			var existing int64
			if err := d.DB.WithContext(c.Request.Context()).Table("user_roles").Count(&existing).Error; err != nil {
				c.JSON(500, gin.H{"error": "internal_error"})
				return
			}
			if existing > 0 {
				c.JSON(409, gin.H{"error": "bootstrap_already_completed"})
				return
			}

			userID, _ := c.Get(platformMiddleware.ContextUserIDKey)
			uid, _ := userID.(string)
			if uid == "" {
				c.JSON(401, gin.H{"error": "invalid_user_context"})
				return
			}

			if err := d.RBACService.AssignRoleToUser(c.Request.Context(), uid, "00000000-0000-0000-0000-000000000001"); err != nil {
				c.JSON(500, gin.H{"error": "internal_error"})
				return
			}

			c.JSON(200, gin.H{"ok": true})
		})
	}
	return nil
}

func (rbacMod) AutoMigrate(db *gorm.DB) error {
	return rbacp.AutoMigrate(db)
}

func (rbacMod) Seed(db *gorm.DB, _ kit.SeedOpts) error {
	return rbacseed.SeedDemo(db)
}

func (rbacMod) Start(context.Context, *kit.Deps) error { return nil }

func (rbacMod) Permissions() []string {
	return []string{
		"admin:ping",
		"rbac:manage",
	}
}

// Module is the RBAC slice.
var Module kit.Module = rbacMod{}
