package main

import (
	"hexad/internal/adapter/database/mongodb"
	"hexad/internal/adapter/handler/http"
	"hexad/internal/core/service"
)

func main() {
	userRepository := mongodb.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	httpServer := http.NewServer("3000", userHandler)
	httpServer.Run()
}
