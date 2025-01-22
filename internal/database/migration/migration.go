package main

import (
	"fmt"
	"log"
	"vibe-user/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationScriptPath = "file://./internal/database/migration/script"

func main() {
	cfg := config.LoadConfig()
	m, err := migrate.New(migrationScriptPath, fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
