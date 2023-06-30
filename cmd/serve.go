package cmd

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal/controllers"
	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var (
	db database.Db = database.Db{}
)

type Serve struct{}

func (s Serve) Execute() error {
	if err := database.Context.Connect(); err != nil {
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

	r.Get("/api/games", controllers.GameController{}.Index)
	r.Get("/api/games/{id}", controllers.GameController{}.Show)
	r.Get("/api/games/{id}/similar", controllers.GameController{}.Similar)
	r.Get("/api/developers/{id}/games", controllers.DeveloperController{}.Games)
	r.Get("/api/genres/{id}/games", controllers.GenreController{}.Games)

	return http.ListenAndServe(":80", r)
}
