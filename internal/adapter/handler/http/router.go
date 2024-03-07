package http

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	userHandler *UserHandler
}

func newRouter(
	userHandler *UserHandler,
) Router {
	return Router{
		userHandler: userHandler,
	}
}

func (r *Router) serve(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	r.userV1Router(v1)
}
