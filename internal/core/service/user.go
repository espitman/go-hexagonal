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
	var resp []domain.User
	for _, user := range users {
		resp = append(resp, domain.User{
			FName: user.FName,
			LName: user.LName,
		})
	}
	return resp, nil
}
