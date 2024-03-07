package http

import (
	"github.com/gofiber/fiber/v2"
	"hexad/internal/core/port"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	resp, _ := h.userService.List(c, 0, 10)
	return c.JSON(resp)
}
