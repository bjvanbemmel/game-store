package migrations

import (
	"github.com/bjvanbemmel/game-store/internal/database"
)

type AddGenreTable struct{}

func (m AddGenreTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS genres (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL
        );
    `)

	return err
}
