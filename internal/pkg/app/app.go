package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"middleware/internal/app/endpoint"
	"middleware/internal/app/mw"
	"middleware/internal/app/service"
)

type App struct {
	edp  *endpoint.Endpoint
	svc  *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	app := &App{}
	app.svc = service.New()
	app.edp = endpoint.New(app.svc)

	app.echo = echo.New()

	app.echo.Use(mw.RoleCheck)

	app.echo.GET("/status", app.edp.Status)

	return app, nil
}

func (a *App) Run() error {
	fmt.Println("server running")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
