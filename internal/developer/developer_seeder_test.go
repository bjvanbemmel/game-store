package developer

import (
	"testing"

	"github.com/bjvanbemmel/game-store/internal"
)

func TestDeveloperSeederSuccessfullySeeds(t *testing.T) {
	var db internal.Db = internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator DeveloperMigrator = DeveloperMigrator{}
	var seeder DeveloperSeeder = DeveloperSeeder{}

	if err := migrator.Migrate(db.DB); err != nil {
		t.Fatal(err)
	}

	if err := seeder.Seed(db.DB); err != nil {
		t.Fatal(err)
	}
}
