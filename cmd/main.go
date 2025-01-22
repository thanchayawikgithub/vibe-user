package main

import (
	"vibe-user/internal/config"
	"vibe-user/internal/database"
	"vibe-user/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(&cfg.Database)

	server.Start(cfg, db)
}
