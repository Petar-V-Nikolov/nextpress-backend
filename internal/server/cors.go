package server

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
)

// buildCORSConfig returns gin-contrib/cors settings from the environment.
// If CORS_ORIGINS is set (comma-separated), only those origins are allowed
// and credentials are supported. Otherwise any origin is allowed with wildcard
// request headers, which browsers treat as incompatible with credentials.
func buildCORSConfig() cors.Config {
	origins := getCORSOrigins()
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	if len(origins) > 0 {
		return cors.Config{
			AllowOrigins:     origins,
			AllowMethods:     methods,
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Length"},
		}
	}
	return cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    methods,
		AllowHeaders:    []string{"*"},
		ExposeHeaders:   []string{"*"},
	}
}

func getCORSOrigins() []string {
	s := strings.TrimSpace(os.Getenv("CORS_ORIGINS"))
	if s == "" {
		return nil
	}
	var out []string
	for _, o := range strings.Split(s, ",") {
		if o = strings.TrimSpace(o); o != "" {
			out = append(out, o)
		}
	}
	return out
}
