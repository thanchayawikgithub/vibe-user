package server

import (
	"fmt"
	"log"
	"vibe-user/internal/config"
	"vibe-user/internal/server/middleware"
	"vibe-user/internal/server/router"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

func Start(cfg *config.Config, db *gorm.DB) {
	e := echo.New()

	mw := middleware.NewMiddleware(e)
	mw.RegisterMiddlewares()

	router := router.NewRouter(e, db, cfg)
	router.RegisterRoutes()

	log.Fatal(e.Start(fmt.Sprintf(":%d", cfg.App.Port)))
}
