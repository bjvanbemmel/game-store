package game

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
)

type GameController struct {
	Db *internal.Db

	internal.Controller
}

func (c GameController) Index(w http.ResponseWriter, r *http.Request) {
	var games []*Game = []*Game{}

	if err := c.Db.Ping(); err != nil {
		c.ResponseError(w, err)
	}

	if err := Paginate(c.Db.DB, 1, 10, &games); err != nil {
		c.ResponseError(w, err)
	}

	c.Response(w, games)
}
