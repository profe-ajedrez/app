package main

import (
	"net/http"
	"testing"

	"github.com/profe-ajedrez/app/chiapp"
)

func BenchmarkSetup(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		app := chiapp.New()
		app.WithHandlers(handlersChi)
	}
}

func BenchmarkExecution(b *testing.B) {

	for i := 0; i <= b.N; i++ {
		_, err := http.Get("http://localhost:3000/clients.json")

		if err != nil {
			b.Log(err.Error())
			b.FailNow()
		}

	}
}
