package game

import (
	"database/sql"
	"testing"
)

func TestGameSeederSuccessfullySeeds(t *testing.T) {
	var dsn string = "file:unit_tests.db?cache=shared&mode=memory"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	var migrator GameMigrator = GameMigrator{}
	var seeder GameSeeder = GameSeeder{}

	if err := migrator.Migrate(db); err != nil {
		t.Fatal(err)
	}

	if err := seeder.Seed(db); err != nil {
		t.Fatal(err)
	}
}
