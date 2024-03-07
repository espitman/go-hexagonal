package port

import (
	"github.com/gofiber/fiber/v2"
	"hexad/internal/core/domain"
)

/**
 * UserService implemented by service.UserService interface
 */

type UserService interface {
	List(ctx *fiber.Ctx, skip, limit uint64) ([]domain.User, error)
}

/**
 * UserRepository implemented by repository.UserRepository interface
 */

type UserRepository interface {
	List(ctx *fiber.Ctx, skip, limit uint64) ([]domain.User, error)
}
