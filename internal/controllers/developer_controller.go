package controllers

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/bjvanbemmel/game-store/internal/dto"
	"github.com/go-chi/chi/v5"
)

type DeveloperController struct {
	internal.Controller
}

func (c DeveloperController) Games(w http.ResponseWriter, r *http.Request) {
	var games []dto.Game = []dto.Game{}
	var developer dto.Developer

	if chi.URLParam(r, "id") == "" {
		c.ResponseError(w, internal.ErrEmptyParameter)
		return
	}

	row := database.Context.QueryRow(`
        SELECT * FROM developers WHERE id = $1;
    `, chi.URLParam(r, "id"))
	row.Scan(&developer.Id, &developer.Name)

	rows, err := database.Context.Query(`
        SELECT g.* FROM games g
        JOIN game_developer gd
        ON g.id = gd.game_id
        JOIN developers d
        ON gd.developer_id = d.id
        WHERE d.id = $1;
    `, chi.URLParam(r, "id"))
	defer rows.Close()

	if err != nil {
		c.ResponseError(w, err)
		return
	}

	for rows.Next() {
		var game dto.Game

		if err := rows.Scan(&game.Id, &game.Title, &game.Price, &game.Thumbnail, &game.Description, &game.ReleaseDate, &game.Rating); err != nil {
			c.ResponseError(w, err)
			return
		}

		if _, err := game.FullFetch(); err != nil {
			c.ResponseError(w, err)
			return
		}

		games = append(games, game)
	}

	c.Response(w, map[string]any{
		"developer": developer,
		"games":     games,
	})
}
