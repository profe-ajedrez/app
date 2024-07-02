package middleware

import (
	"log"
	"net/http"
)

func ExampleMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("running before Example Middleware")
	next(rw, r)
	log.Println("running after Example  Middleware")
}
