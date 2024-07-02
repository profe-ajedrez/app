package fiberapp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
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

	a.handlers(app)

	return adaptor.FiberApp(app)

}
