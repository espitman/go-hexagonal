package main

import (
	_ "github.com/espitman/go-hexagonal/cmd/http/docs"
	"github.com/espitman/go-hexagonal/internal/adapter/database/postgre"
	"github.com/espitman/go-hexagonal/internal/adapter/handler/http"
	"github.com/espitman/go-hexagonal/internal/core/service"
	"github.com/espitman/go-hexagonal/internal/core/util"
	superConf "github.com/espitman/go-super-conf"
)

// @title           Salvation2 - price service
// @version         1.0
// @description     This is jabama Price service
// @contact.name   API Support
// @contact.email  s.heidar@jabama.com
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	validate := util.NewValidator()

	pgDB := postgre.NewDB(superConf.Get("db.postgres.dsn"))

	calendarRepository := postgre.NewCalendarRepository(pgDB.Connection)
	calendarService := service.NewCalendarService(calendarRepository)
	calendarHandler := http.NewCalendarHandler(validate, calendarService)

	httpServer := http.NewServer(superConf.Get("app.port"), calendarHandler)
	httpServer.Run()
}
