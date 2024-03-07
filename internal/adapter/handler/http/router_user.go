package http

import "github.com/gofiber/fiber/v2"

func (r *Router) userV1Router(v1 fiber.Router) {
	v1User := v1.Group("/user")
	v1User.Get("/list", r.userHandler.List)
}
