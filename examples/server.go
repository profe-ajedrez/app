package examples

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const Timeout = 30 * time.Second

type ServerOption func(server *http.Server)

func Start(port string, handler http.Handler, options ...ServerOption) error {

	srv := &http.Server{
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
		Addr:         ":" + port,
		Handler:      handler,
	}

	for _, o := range options {
		o(srv)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		log.Println("Stopping server")
		err := srv.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	log.Println(fmt.Sprintf("Service listening on port %s", port))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// WithReadTimeout configure http.Server parameter ReadTimeout
func WithReadTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.ReadTimeout = t
	}
}

// WithWriteTimeout configure http.Server parameter WriteTimeout
func WithWriteTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.WriteTimeout = t
	}
}
