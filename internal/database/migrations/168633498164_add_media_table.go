package migrations

import "github.com/bjvanbemmel/game-store/internal/database"

type AddMediaTable struct{}

func (m AddMediaTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS media (
            id SERIAL PRIMARY KEY,
            game_id INTEGER REFERENCES games(id),
            type INTEGER NOT NULL,
            uri VARCHAR(355) NOT NULL
        )
    `)

	return err
}
