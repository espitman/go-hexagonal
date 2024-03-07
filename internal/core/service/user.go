package service

import (
	"github.com/gofiber/fiber/v2"
	"hexad/internal/adapter/database/mongodb/repository"
	"hexad/internal/core/domain"
)

/**
 * UserService implements port.UserService interface
 */

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) List(ctx *fiber.Ctx, skip, limit uint64) ([]domain.User, error) {
	return s.UserRepository.List(ctx, skip, limit)
}
