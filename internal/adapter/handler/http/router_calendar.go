package http

import "github.com/gofiber/fiber/v2"

func (r *Router) calendarV1Router(v1 fiber.Router) {
	v1Calendar := v1.Group("/calendar")

	v1Calendar.Use(AuthMiddleware)

	v1Calendar.Post("/", r.calendarHandler.Create)
	v1Calendar.Get("/", r.calendarHandler.Get)
}
