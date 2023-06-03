package migrations

import (
	"github.com/bjvanbemmel/game-store/internal/database"
)

type AddDevelopersTable struct{}

func (m AddDevelopersTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS developers (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL
        );
    `)

	return err
}
