package migrations

import (
	"github.com/bjvanbemmel/game-store/internal/database"
)

type AddGameDeveloperPivotTable struct{}

func (m AddGameDeveloperPivotTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS game_developer (
            game_id INTEGER REFERENCES games(id),
            developer_id INTEGER REFERENCES developers(id)
        );
    `)

	return err
}
