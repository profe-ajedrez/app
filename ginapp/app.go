package ginapp

import (
	"github.com/gin-gonic/gin"
	"github.com/profe-ajedrez/app/errs"
)

type GinHandlers func(r *gin.Engine)

type App struct {
	handlers GinHandlers
}

func New() *App {
	return &App{}
}

func (a *App) WithHandlers(handlers GinHandlers) {
	a.handlers = handlers
}

func (a App) Handlers() *gin.Engine {
	r := gin.Default()

	if a.handlers == nil {
		panic(errs.NewGinAppErr("there is no handlers defined"))
	}

	a.handlers(r)

	return r
}
