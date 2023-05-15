package game

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGameMigratorSuccessfullyMigrates(t *testing.T) {
	var dsn string = "file:unit_tests.db?cache=shared&mode=memory"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	var migrator GameMigrator = GameMigrator{}

	if err := migrator.Clear(db); err != nil {
		t.Fatal(err)
	}

	if err := migrator.Migrate(db); err != nil {
		t.Fatal(err)
	}
}
