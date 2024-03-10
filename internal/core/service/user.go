package service

import (
	"github.com/gofiber/fiber/v2"
	"hexad/internal/core/domain"
	"hexad/internal/core/port"
)

/**
 * UserService implements port.UserService interface
 */

type UserService struct {
	UserRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) List(ctx *fiber.Ctx, skip, limit uint64) ([]domain.User, error) {
	users, _ := s.UserRepository.List(ctx, skip, limit)
	return users, nil
}
