package cmd

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	db internal.Db = internal.Db{}
)

type Serve struct{}

func (s Serve) Execute() error {
	if err := db.Connect(); err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"http://*"},
		},
	))

	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})

	return http.ListenAndServe(":80", r)
}
