package router

import (
	"vibe-user/internal/modules/auth/handler"
	"vibe-user/internal/modules/user/repository"
	"vibe-user/internal/modules/user/service"
)

func (r *Router) RegisterAuthRoutes() {
	userRepository := repository.NewUserRepository(r.db)
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(userService, &r.cfg.Oauth)
	auth := r.server.Group("/auth")
	auth.GET("/google/login", authHandler.GoogleLogin)
	auth.GET("/google/callback", authHandler.GoogleCallback)
}
