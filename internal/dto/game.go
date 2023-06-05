package dto

import "github.com/bjvanbemmel/game-store/internal/database"

type Game struct {
	Id         int
	Title      string
	Price      float32
	Developers []Developer
	Categories []Category
}

func (g *Game) FetchDevelopers() (*Game, error) {
	rows, err := database.Context.Query(`
        SELECT * FROM game_developer gd
        JOIN developer d
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
