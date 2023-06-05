package controllers

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/bjvanbemmel/game-store/internal/dto"
)

type GameController struct {
	internal.Controller
}

func (c GameController) Index(w http.ResponseWriter, r *http.Request) {
	var games []dto.Game

	rows, err := database.Context.Query(`
        SELECT * FROM games g LIMIT 15;
    `)
	// defer rows.Close()

	if err != nil {
		c.ResponseError(w, err)
		return
	}

	for rows.Next() {
		var game dto.Game

		if err := rows.Scan(&game.Id, &game.Title, &game.Price, &game.Thumbnail, &game.Description); err != nil {
			c.ResponseError(w, err)
			return
		}

		if _, err := game.FullFetch(); err != nil {
			c.ResponseError(w, err)
			return
		}

		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		c.ResponseError(w, err)
		return
	}

	c.Response(w, games)
}
