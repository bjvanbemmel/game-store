package controllers

import (
	"net/http"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/bjvanbemmel/game-store/internal/dto"
	"github.com/go-chi/chi/v5"
)

type GameController struct {
	internal.Controller
}

func (c GameController) Index(w http.ResponseWriter, r *http.Request) {
	var games []dto.Game = []dto.Game{}

	if r.URL.Query().Get("keyword") != "" {
		c.PartialSearch(w, r)
		return
	}

	rows, err := database.Context.Query(`
        SELECT * FROM games LIMIT 15;
    `)
	defer rows.Close()

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

func (c GameController) Show(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "id") == "" {
		c.ResponseError(w, internal.ErrEmptyParameter)
		return
	}

	row := database.Context.QueryRow(`
        SELECT * FROM games WHERE id = $1
    `, chi.URLParam(r, "id"))

	var game dto.Game
	if err := row.Scan(&game.Id, &game.Title, &game.Price, &game.Thumbnail, &game.Description); err != nil {
		c.ResponseError(w, err)
		return
	}

	if _, err := game.FullFetch(); err != nil {
		c.ResponseError(w, err)
		return
	}

	c.Response(w, game)
}

func (c GameController) PartialSearch(w http.ResponseWriter, r *http.Request) {
	var games []dto.Game = []dto.Game{}

	rows, err := database.Context.Query(`
        SELECT * FROM games WHERE title ILIKE '%' || $1 || '%';
    `, r.URL.Query().Get("keyword"))

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

	c.Response(w, games)
}

func (c GameController) Similar(w http.ResponseWriter, r *http.Request) {
	var games []dto.Game
	var curGame dto.Game

	if chi.URLParam(r, "id") == "" {
		c.ResponseError(w, internal.ErrEmptyParameter)
		return
	}

	row := database.Context.QueryRow(`
        SELECT * FROM games WHERE id = $1
    `, chi.URLParam(r, "id"))

	if err := row.Scan(&curGame.Id, &curGame.Title, &curGame.Price, &curGame.Thumbnail, &curGame.Description); err != nil {
		c.ResponseError(w, err)
		return
	}

	if _, err := curGame.FetchGenres(); err != nil {
		c.ResponseError(w, err)
		return
	}

	games, err := curGame.Similar()
	if err != nil {
		c.ResponseError(w, err)
		return
	}

	c.Response(w, games)
}
