package http

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiberRouter     fiber.Router
	calendarHandler *CalendarHandler
}

func newRouter(
	calendarHandler *CalendarHandler,
) Router {
	return Router{
		calendarHandler: calendarHandler,
	}
}

func (r *Router) serve(app *fiber.App) {
	r.fiberRouter = app.Group("/")
	api := app.Group("/api")
	v1 := api.Group("/v1")

	r.swaggerRouter()
	r.calendarV1Router(v1)
}
