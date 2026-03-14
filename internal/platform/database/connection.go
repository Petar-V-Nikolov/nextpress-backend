package database

import (
	"fmt"

	"gorm.io/gorm"
)

// Config holds database configuration for connection.
type Config struct {
	Driver   string // Driver name (postgres, mysql)
	Host     string
	Port     string
	User     string
	Password string
	Name     string // Database name
	SSLMode  string // SSL mode (postgres-specific)
}

// New returns a new *gorm.DB instance based on the driver specified in Config.
// This function allows the application to remain database-agnostic.
func New(cfg Config) (*gorm.DB, error) {
	switch cfg.Driver {
	case "postgres":
		return newPostgres(cfg)
	// case "mysql":
	//  return newMySQL(cfg)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}
}
