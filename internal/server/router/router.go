package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server *echo.Echo
}

func NewRouter(server *echo.Echo) *Router {
	return &Router{server}
}

func (r *Router) Register() {
	r.server.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
}
