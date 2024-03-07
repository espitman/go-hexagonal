package http

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port        string
	UserHandler *UserHandler
}

func NewHttpServer(
	port string,
	userHandler *UserHandler,
) Server {
	return Server{
		Port:        port,
		UserHandler: userHandler,
	}
}

func (s *Server) Run() {
	app := fiber.New()
	routes := newRouter(s.UserHandler)
	routes.serve(app)
	app.Listen(":" + s.Port)
}
