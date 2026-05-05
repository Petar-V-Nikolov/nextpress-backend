package kit

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// ResolveModulesFromRegistry filters registry by MODULES env (comma-separated ids).
// Empty MODULES uses the full registry slice unchanged.
func ResolveModulesFromRegistry(log *zap.SugaredLogger, registry []Module) []Module {
	raw := strings.TrimSpace(os.Getenv("MODULES"))
	if raw == "" {
		return registry
	}
	requested := map[string]struct{}{}
	for _, p := range strings.Split(raw, ",") {
		id := strings.ToLower(strings.TrimSpace(p))
		if id == "" {
			continue
		}
		requested[id] = struct{}{}
	}
	if len(requested) == 0 {
		return registry
	}
	applyImplicitDeps(requested)
	var out []Module
	for _, m := range registry {
		if _, ok := requested[m.ID()]; ok {
			out = append(out, m)
		}
	}
	for id := range requested {
		found := false
		for _, m := range registry {
			if m.ID() == id {
				found = true
				break
			}
		}
		if !found && log != nil {
			log.Warnw("unknown module id in MODULES env (ignored)", "module", id)
		}
	}
	return out
}

func applyImplicitDeps(set map[string]struct{}) {
	if _, ok := set["auth"]; ok {
		set["user"] = struct{}{}
		set["rbac"] = struct{}{}
	}
	content := false
	for _, id := range []string{"posts", "pages", "taxonomy", "media"} {
		if _, ok := set[id]; ok {
			content = true
			break
		}
	}
	if content {
		set["user"] = struct{}{}
		set["rbac"] = struct{}{}
		set["auth"] = struct{}{}
	}
	if _, ok := set["posts"]; ok {
		set["taxonomy"] = struct{}{}
		set["media"] = struct{}{}
	}
}

// CollectPermissionCodes merges module-declared default permission codes (deduped).
func CollectPermissionCodes(modules []Module) []string {
	seen := map[string]struct{}{}
	var out []string
	for _, m := range modules {
		for _, c := range m.Permissions() {
			c = strings.TrimSpace(c)
			if c == "" {
				continue
			}
			if _, ok := seen[c]; ok {
				continue
			}
			seen[c] = struct{}{}
			out = append(out, c)
		}
	}
	return out
}
