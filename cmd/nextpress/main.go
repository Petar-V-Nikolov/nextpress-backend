package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Petar-V-Nikolov/nextpress-backend/internal/config"
	platformLogger "github.com/Petar-V-Nikolov/nextpress-backend/internal/platform/logger"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/server"
)

func main() {
	// Use a dedicated context for the lifetime of the application; this makes it
	// straightforward to propagate graceful shutdown signals to all subsystems.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize structured logging as early as possible so we can rely on a
	// consistent, production-ready logger for all subsequent operations.
	baseLogger, err := platformLogger.New()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	logger := baseLogger.Sugar()
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(baseLogger)

	logger.Infow("starting nextpress-backend",
		"version", "0.1.0-phase1",
	)

	// Load environment variables (from .env if present) and app configuration
	// before touching any external resources (DB, message buses, etc.).
	config.LoadEnv()
	appCfg := config.LoadAppConfig()
	dbCfg := config.LoadDBConfig()

	// Use Gin as the central HTTP router; we keep the setup centralized in the
	// server package so that future modules can register routes cleanly.
	engine := gin.New()
	server.ConfigureEngine(engine, logger)

	// Phase 1: We bootstrap the HTTP server only. The global DB instance and
	// module wiring will be layered on top in later phases, but we already
	// pass configuration through the Server to keep the surface stable.
	srv := server.NewServer(engine, appCfg, dbCfg, logger)

	// Run the HTTP server in its own goroutine so that we can listen for OS
	// signals and coordinate a controlled shutdown sequence.
	go func() {
		if err := srv.Start(); err != nil {
			logger.Fatalw("http server exited with error",
				"error", err,
			)
		}
	}()

	// Capture SIGINT/SIGTERM and use them as a trigger for graceful shutdown.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigCh
	logger.Infow("received shutdown signal",
		"signal", sig.String(),
	)

	// Apply a hard timeout to shutdown to avoid hanging the process indefinitely.
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		// At this point we log to the standard log package as a last resort in case
		// the structured logger is already partially torn down.
		log.Printf("graceful shutdown failed: %v\n", err)
	}

	logger.Info("nextpress-backend stopped cleanly")
}

