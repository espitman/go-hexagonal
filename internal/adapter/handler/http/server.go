package http

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port            string
	CalendarHandler *CalendarHandler
}

func NewServer(
	port string,
	userHandler *CalendarHandler,
) Server {
	return Server{
		Port:            port,
		CalendarHandler: userHandler,
	}
}

func (s *Server) Run() {
	app := fiber.New()
	routes := newRouter(s.CalendarHandler)
	routes.serve(app)
	_ = app.Listen(":" + s.Port)
}
