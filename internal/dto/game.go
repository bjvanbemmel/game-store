package dto

import (
	"github.com/bjvanbemmel/game-store/internal/database"
)

type Game struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Price       float32     `json:"price"`
	Thumbnail   string      `json:"thumbnail"`
	Description string      `json:"description"`
	Developers  []Developer `json:"developers"`
	Genres      []Genre     `json:"genres"`
}

func (g *Game) FullFetch() (*Game, error) {
	if _, err := g.FetchDevelopers(); err != nil {
		return nil, err
	}

	if _, err := g.FetchGenres(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) FetchDevelopers() (*Game, error) {
	rows, err := database.Context.Query(`
        SELECT d.id, d.name FROM game_developer gd
        JOIN developers d
        ON gd.developer_id = d.id
        WHERE gd.game_id = $1;
    `, g.Id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dev Developer

		if err := rows.Scan(&dev.Id, &dev.Name); err != nil {
			return nil, err
		}

		g.Developers = append(g.Developers, dev)
	}

	return g, nil
}

func (g *Game) FetchGenres() (*Game, error) {
	rows, err := database.Context.Query(`
        SELECT ge.id, ge.name FROM game_genre gg
        JOIN genres ge
        ON gg.genre_id = ge.id
        WHERE gg.game_id = $1;
    `, g.Id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var gen Genre

		if err := rows.Scan(&gen.Id, &gen.Name); err != nil {
			return nil, err
		}

		g.Genres = append(g.Genres, gen)
	}

	return g, nil
}
