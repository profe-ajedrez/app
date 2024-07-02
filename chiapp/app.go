package chiapp

import (
	"github.com/go-chi/chi/v5"
	"github.com/profe-ajedrez/app/errs"
)

type ChiEngine func() *chi.Mux
type ChiHandlers func() *chi.Mux

type App struct {
	handlers ChiHandlers
}

func New() *App {
	return &App{}
}

func (a *App) WithHandlers(handlers ChiHandlers) {
	a.handlers = handlers
}

func (a App) Engine() *chi.Mux {
	return chi.NewRouter()
}

func (a App) Handlers() *chi.Mux {
	if a.handlers == nil {
		panic(errs.NewChiAppErr("there is no handlers defined"))
	}

	return a.handlers()
}
