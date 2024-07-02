package chiapp

import (
	"github.com/go-chi/chi/v5"
	"github.com/profe-ajedrez/app/errs"
)

type ChiHandlers func(*chi.Mux)

type App struct {
	handlers ChiHandlers
}

func New() *App {
	return &App{}
}

func (a *App) WithHandlers(handlers ChiHandlers) {
	a.handlers = handlers
}

func (a App) Handlers() *chi.Mux {
	r := chi.NewRouter()

	if a.handlers == nil {
		panic(errs.NewChiAppErr("there is no handlers defined"))
	}

	a.handlers(r)

	return r
}
