package controllers

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/bjvanbemmel/game-store/internal/dto"
	"github.com/go-chi/chi/v5"
)

type GenreController struct {
	internal.Controller
}

func (c GenreController) Games(w http.ResponseWriter, r *http.Request) {
	var games []dto.Game = []dto.Game{}
	var genre dto.Genre

	var id string = chi.URLParam(r, "id")
	if id == "" {
		c.ResponseError(w, internal.ErrEmptyParameter)
		return
	}

	row := database.Context.QueryRow(`
        SELECT * FROM genres WHERE id = $1;
    `, id)
	row.Scan(&genre.Id, &genre.Name)

	rows, err := database.Context.Query(`
        SELECT g.* FROM games g
        JOIN game_genre gg
        ON gg.game_id = g.id
        JOIN genres ge
        ON ge.id = gg.genre_id
        WHERE ge.id = $1;
    `, id)
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
		"genre": genre,
		"games": games,
	})
}
