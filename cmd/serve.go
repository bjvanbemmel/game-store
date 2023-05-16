package cmd

import (
	"fmt"
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/game"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	db             internal.Db         = internal.Db{}
	gameController game.GameController = game.GameController{
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

	fmt.Println(db.Ping())

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})

	r.Get("/api/games", gameController.Index)

	return http.ListenAndServe(":80", r)
}
