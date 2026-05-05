// Package appregistry is the compile-time default module list (migrations + HTTP + seed).
// It lives outside internal/app and internal/kit so cmd/* and pkg/seed can import it without cycles.
package appregistry

import (
	"github.com/nextpresskit/backend/internal/kit"
	authmod "github.com/nextpresskit/backend/internal/modules/auth/module"
	mediamod "github.com/nextpresskit/backend/internal/modules/media/module"
	pagesmod "github.com/nextpresskit/backend/internal/modules/pages/module"
	postsmod "github.com/nextpresskit/backend/internal/modules/posts/module"
	rbacmod "github.com/nextpresskit/backend/internal/modules/rbac/module"
	taxonomymod "github.com/nextpresskit/backend/internal/modules/taxonomy/module"
	usermod "github.com/nextpresskit/backend/internal/modules/user/module"
)

// ModuleRegistry is the FK-safe default module order.
func ModuleRegistry() []kit.Module {
	return []kit.Module{
		usermod.Module,
		rbacmod.Module,
		authmod.Module,
		taxonomymod.Module,
		mediamod.Module,
		postsmod.Module,
		pagesmod.Module,
	}
}
