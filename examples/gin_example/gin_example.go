package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/profe-ajedrez/app/examples"
	"github.com/profe-ajedrez/app/ginapp"
	"github.com/profe-ajedrez/app/services"
)

func main() {
	app := ginapp.New()

	app.WithHandlers(handlersGin)

	log.Println("Starting server")
	err := examples.Start("3000", app.Handlers())

	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}

func handlersGin(r *gin.Engine) {
	r.Handle("GET", "/clients.json", ClientHandler())
	r.Handle("GET", "/health.json", ginHealthHandler)
}

func ClientHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := services.GetClients(c)
		c.JSON(http.StatusOK, response)

	}
}

func ginHealthHandler(c *gin.Context) {
	response := struct {
		Health string `json:"health"`
	}{
		Health: "OK",
	}
	c.JSON(http.StatusOK, response)
}
