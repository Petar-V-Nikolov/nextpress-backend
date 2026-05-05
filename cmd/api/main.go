package main

import (
	"context"
	"log"

	"github.com/nextpresskit/backend/internal/app"
	"github.com/nextpresskit/backend/internal/appregistry"
	"github.com/nextpresskit/backend/internal/config"
	"github.com/nextpresskit/backend/internal/kit"
	platformLogger "github.com/nextpresskit/backend/internal/platform/logger"
)

var version = "dev"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config.LoadEnv()

	baseLogger, err := platformLogger.New()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	mods := kit.ResolveModulesFromRegistry(baseLogger.Sugar(), appregistry.ModuleRegistry())
	if err := app.Run(ctx, version, baseLogger, mods); err != nil {
		log.Fatalf("run: %v", err)
	}
}
