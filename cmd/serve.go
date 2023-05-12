package cmd

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Serve struct{}

func (s Serve) Execute() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})

	return http.ListenAndServe(":8123", r)
}
