package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
	server *echo.Echo
}

func NewMiddleware(server *echo.Echo) *Middleware {
	return &Middleware{server}
}

func (m *Middleware) Register() {
	m.server.Use(middleware.Logger())
	m.server.Use(middleware.Recover())
}
