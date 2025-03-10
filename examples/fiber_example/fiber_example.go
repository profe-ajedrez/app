package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/profe-ajedrez/app/examples"
	"github.com/profe-ajedrez/app/fiberapp"
	"github.com/profe-ajedrez/app/services"
)

func main() {
	app := fiberapp.New()

	app.WithHandlers(handlersFiber)

	log.Println("Starting server")
	err := examples.Start("3000", app.Handlers())

	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}

func handlersFiber(r *fiber.App) {
	r.Get("/clients.json", ClientHandler())
	r.Get("/health.json", ginHealthHandler)
}

func ClientHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		response := services.GetClients(c.Context())
		return c.Status(http.StatusOK).JSON(response)
	}
}

func ginHealthHandler(c *fiber.Ctx) error {
	response := struct {
		Health string `json:"health"`
	}{
		Health: "OK",
	}

	return c.Status(http.StatusOK).JSON(response)
}
