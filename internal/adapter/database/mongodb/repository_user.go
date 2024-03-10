package mongodb

import (
	"github.com/gofiber/fiber/v2"
)

type UserRepository struct {
}

/**
 * UserRepository implements port.UserRepository interface
 */

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (s *UserRepository) List(ctx *fiber.Ctx, skip, limit uint64) ([]UserSchema, error) {
	var resp []UserSchema
	resp = append(resp, UserSchema{
		FName: "Saeed",
		LName: "Heidari",
	}, UserSchema{
		FName: "John",
		LName: "Smith",
	})
	return resp, nil
}
