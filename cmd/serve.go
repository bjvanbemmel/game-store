package cmd

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/developer"
	"github.com/bjvanbemmel/game-store/internal/game"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	db             internal.Db         = internal.Db{}
	gameController game.GameController = game.GameController{
		Db: &db,
	}
	developerController developer.DeveloperController = developer.DeveloperController{
		Db: &db,
	}
)

type Serve struct{}

func (s Serve) Execute() error {
	if err := db.Environment("prod"); err != nil {
		return err
	}

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

	r.Get("/api/games", gameController.Index)

	r.Get("/api/developers", developerController.Index)

	return http.ListenAndServe(":80", r)
}
