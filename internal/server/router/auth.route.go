package router

import (
	"vibe-user/internal/modules/auth/handler"
	"vibe-user/internal/modules/auth/service"
	"vibe-user/internal/modules/user/repository"
)

func (r *Router) RegisterAuthRoutes() {
	userRepository := repository.NewUserRepository(r.db)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService, &r.cfg.Oauth)
	auth := r.server.Group("/auth")
	auth.GET("/google/login", authHandler.GoogleLogin)
	auth.GET("/google/callback", authHandler.GoogleCallback)
}
