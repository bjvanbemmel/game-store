package developer

import (
	"testing"

	"github.com/bjvanbemmel/game-store/internal"
)

func TestDeveloperMigratorSuccessfullyMigrates(t *testing.T) {
	var db internal.Db = internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator DeveloperMigrator = DeveloperMigrator{}

	if err := migrator.Clear(db.DB); err != nil {
		t.Fatal(err)
	}

	if err := migrator.Migrate(db.DB); err != nil {
		t.Fatal(err)
	}
}
