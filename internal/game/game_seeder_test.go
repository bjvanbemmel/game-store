package game

import (
	"testing"

	"github.com/bjvanbemmel/game-store/internal"
)

func TestGameSeederSuccessfullySeeds(t *testing.T) {
	var db internal.Db = internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator GameMigrator = GameMigrator{}
	var seeder GameSeeder = GameSeeder{}

	if err := migrator.Migrate(db.DB); err != nil {
		t.Fatal(err)
	}

	if err := seeder.Seed(db.DB); err != nil {
		t.Fatal(err)
	}
}
