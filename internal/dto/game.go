package dto

import (
	"time"

	"github.com/bjvanbemmel/game-store/internal/database"
)

type Game struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	Price       float32       `json:"price"`
	Thumbnail   string        `json:"thumbnail"`
	Description string        `json:"description"`
	Media       []Media       `json:"media"`
	Developers  []Developer   `json:"developers"`
	Genres      []Genre       `json:"genres"`
	PlayerCount []PlayerCount `json:"player_count"`
	ReleaseDate time.Time     `json:"release_date"`
}

func (g *Game) FullFetch() (*Game, error) {
	if _, err := g.FetchDevelopers(); err != nil {
		return nil, err
	}

	if _, err := g.FetchGenres(); err != nil {
		return nil, err
	}

	if _, err := g.FetchMedia(); err != nil {
		return nil, err
	}

	if _, err := g.FetchPlayerCounts(); err != nil {
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

func (g *Game) FetchMedia() (*Game, error) {
	rows, err := database.Context.Query(`
        SELECT m.id, m.type, m.uri FROM media m
        JOIN games g
        ON g.id = m.game_id
        WHERE m.game_id = $1;
    `, g.Id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var med Media

		if err := rows.Scan(&med.Id, &med.Type, &med.Uri); err != nil {
			return nil, err
		}

		g.Media = append(g.Media, med)
	}

	return g, nil
}

func (g *Game) FetchPlayerCounts() (*Game, error) {
	rows, err := database.Context.Query(`
        SELECT p.hour, p.count FROM game_hour_players p
        JOIN games g
        ON p.game_id = g.id
        WHERE p.game_id = $1;
    `, g.Id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var count PlayerCount

		if err := rows.Scan(&count.Hour, &count.Count); err != nil {
			return nil, err
		}

		g.PlayerCount = append(g.PlayerCount, count)
	}

	return g, nil
}

func (g *Game) Similar() ([]Game, error) {
	var games []Game

	rows, err := database.Context.Query(`
        SELECT g.* FROM games g
        JOIN game_genre gg
        ON gg.game_id = g.id
        JOIN genres ge
        ON gg.genre_id = ge.id
        WHERE gg.genre_id IN (
	        SELECT genre_id FROM game_genre
            WHERE game_id = $1
        )
        AND gg.game_id != $1
        GROUP BY g.id;
    `, g.Id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var game Game

		if err := rows.Scan(&game.Id, &game.Title, &game.Price, &game.Thumbnail, &game.Description); err != nil {
			return nil, err
		}

		if _, err := game.FullFetch(); err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}
