package router

import (
	"vibe-user/internal/modules/user/handler"
	"vibe-user/internal/modules/user/repository"
	"vibe-user/internal/modules/user/service"
)

const userPath = "/user"

func (r *Router) RegisterUserRoutes() {
	userRepository := repository.NewUserRepository(r.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user := r.server.Group(userPath)
	user.POST("", userHandler.Create)
	user.GET("/:id", userHandler.FindByID)
}
