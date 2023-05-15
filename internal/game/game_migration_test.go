package game

import (
	"database/sql"
	"testing"
)

func TestGameMigratorSuccessfullyMigrates(t *testing.T) {
	var migrator GameMigrator = GameMigrator{}

	err := migrator.Migrate(sql.DB{})
	if err != nil {
		t.Fatal(err)
	}
}
