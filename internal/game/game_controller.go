package game

import (
	"encoding/json"
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
)

type GameController struct {
	Db *internal.Db
}

func (g GameController) Index(w http.ResponseWriter, r *http.Request) {
	var games []*Game = []*Game{}

	err := g.Db.Ping()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	err = Paginate(g.Db.DB, 1, 10, &games)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	raw, err := json.Marshal(games)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.Write(raw)
}
