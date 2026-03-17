package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ConfigureEngine applies global middleware and registers all top-level routes.
// Feature modules will later plug into the provided router via dedicated
// registration functions to keep boundaries clear.
func ConfigureEngine(engine *gin.Engine, log *zap.SugaredLogger) {
	// In production you typically want to disable Gin's debug output and rely
	// on structured logging instead.
	gin.SetMode(gin.ReleaseMode)

	// Global middleware stack. We keep this minimal in Phase 1 and will extend
	// it (e.g. for authentication, request IDs, metrics) in later phases.
	engine.Use(gin.Recovery())
	engine.Use(requestLoggingMiddleware(log))

	// Health and readiness endpoints are intentionally simple and dependency
	// free in Phase 1. Database and external checks will be wired in later.
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	engine.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
	})

	// Root endpoint to confirm that the service is running and identify the
	// backend explicitly (useful in multi-service environments).
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "nextpress-backend",
		})
	})
}

// requestLoggingMiddleware provides concise, structured logging of incoming
// HTTP traffic. It avoids duplicating Gin's own debug logging while still
// emitting meaningful data for production troubleshooting.
func requestLoggingMiddleware(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		status := c.Writer.Status()
		log.Infow("http request completed",
			"method", method,
			"path", path,
			"status", status,
		)
	}
}

