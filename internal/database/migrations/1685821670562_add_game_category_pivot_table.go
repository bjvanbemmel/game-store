package migrations

import "github.com/bjvanbemmel/game-store/internal/database"

type AddGameCategoryPivotTable struct{}

func (m AddGameCategoryPivotTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS game_category (
            game_id INTEGER REFERENCES games(id),
            category_id INTEGER REFERENCES categories(id)
        );
    `)

	return err
}
