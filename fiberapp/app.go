package fiberapp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/profe-ajedrez/app/errs"
)

type FiberHandlers func(r *fiber.App)

type App struct {
	handlers FiberHandlers
}

func New() *App {
	return &App{}
}

func (a *App) WithHandlers(handlers FiberHandlers) {
	a.handlers = handlers
}

func (a App) Handlers() http.HandlerFunc {
	app := fiber.New()

	if a.handlers == nil {
		panic(errs.NewFiberAppErr("there is no handlers defined"))
	}

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Get("/metrics", monitor.New())

	a.handlers(app)

	return adaptor.FiberApp(app)

}
