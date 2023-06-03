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
            price NUMERIC
        );
    `)

	return err
}
