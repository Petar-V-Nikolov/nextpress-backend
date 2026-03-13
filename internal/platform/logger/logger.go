package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates and returns a structured Zap logger.
// This logger uses ISO8601 timestamps, making logs human-readable and easier to parse.
func New() (*zap.Logger, error) {
	// Start with Zap's production configuration.
	cfg := zap.NewProductionConfig()

	// Set timestamp key and format.
	// ISO8601TimeEncoder produces timestamps like "2026-03-14T15:04:05.999Z07:00".
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Optional customization:
	// You can change log level, encoding, or output paths if needed in the future.
	// Example:
	// cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	// cfg.OutputPaths = []string{"stdout", "logs/app.log"}

	// Build the logger from configuration
	logger, err := cfg.Build()
	if err != nil {
		return nil, err // Return error if logger setup fails
	}

	return logger, nil
}
