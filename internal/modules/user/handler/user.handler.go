package handler

import (
	"net/http"
	"vibe-user/internal/modules/user/entity"
	"vibe-user/internal/modules/user/service"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Create(c echo.Context) error
	FindByID(c echo.Context) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
}

func (h *userHandler) Create(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	savedUser, err := h.userService.Create(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, savedUser)
}

func (h *userHandler) FindByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.userService.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
