package router

import (
	"net/http"
	"vibe-user/internal/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Router struct {
	server *echo.Echo
	db     *gorm.DB
	cfg    *config.Config
}

func NewRouter(server *echo.Echo, db *gorm.DB, cfg *config.Config) *Router {
	return &Router{server, db, cfg}
}

func (r *Router) RegisterRoutes() {
	r.server.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	r.RegisterUserRoutes()
	r.RegisterAuthRoutes()
}
