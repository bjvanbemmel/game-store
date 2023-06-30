package migrations

import (
	"github.com/bjvanbemmel/game-store/internal/database"
)

type AddGamesTable struct{}

func (m AddGamesTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS games (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            price NUMERIC NOT NULL DEFAULT 0.00,
            thumbnail VARCHAR(355) NOT NULL,
            description VARCHAR(3200) NOT NULL,
            release_date DATE NOT NULL DEFAULT NOW(),
            rating NUMERIC NOT NULL DEFAULT -1.00
        );
    `)

	return err
}
