package migrations

import (
	"github.com/bjvanbemmel/game-store/internal/database"
)

type AddCategoriesTable struct{}

func (m AddCategoriesTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS categories (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL
        );
    `)

	return err
}
