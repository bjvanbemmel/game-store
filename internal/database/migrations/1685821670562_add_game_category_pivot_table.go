package migrations

import "github.com/bjvanbemmel/game-store/internal/database"

type AddGameGenrePivotTable struct{}

func (m AddGameGenrePivotTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS game_genre (
            game_id INTEGER REFERENCES games(id),
            genre_id INTEGER REFERENCES genres(id)
        );
    `)

	return err
}
