package game

import "database/sql"

type GameMigrator struct{}

func (m GameMigrator) Migrate(db sql.DB) error {
	return nil
}
