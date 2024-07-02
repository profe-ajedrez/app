package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/profe-ajedrez/app/chiapp"
	"github.com/profe-ajedrez/app/examples"
	"github.com/profe-ajedrez/app/services"
)

func main() {
	app := chiapp.New()

	app.WithHandlers(handlersChi)

	log.Println("Starting server")
	err := examples.Start("3000", app.Handlers())

	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}

func handlersChi(r *chi.Mux) {
	r.Get("/clients.json", ClientHandler())
	r.Get("/health.json", chiHealthHandler)
}

func ClientHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()
		response := services.GetClients(c)

		rp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(rp)
	}
}

func chiHealthHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Health string `json:"health"`
	}{
		Health: "OK",
	}

	rp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rp)
}
