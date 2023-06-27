package migrations

import "github.com/bjvanbemmel/game-store/internal/database"

type AddGameHourPlayersTable struct{}

func (m AddGameHourPlayersTable) Migrate() error {
	_, err := database.Context.Exec(`
        CREATE TABLE IF NOT EXISTS game_hour_players (
            id SERIAL PRIMARY KEY,
            game_id INTEGER REFERENCES games(id),
            count INTEGER DEFAULT 0
        )
    `)

	return err
}
