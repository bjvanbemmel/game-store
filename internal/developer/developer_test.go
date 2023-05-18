package developer

import (
	"testing"

	"github.com/bjvanbemmel/game-store/internal"
)

func TestDeveloperCanInsertAndFindResult(t *testing.T) {
	var db internal.Db = internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator DeveloperMigrator = DeveloperMigrator{}
	migrator.Clear(db.DB)
	migrator.Migrate(db.DB)

	var dev Developer = Developer{
		Name: "Rockstar Games",
	}

	if err := Insert(db.DB, &dev); err != nil {
		t.Fatal(err)
	}

	var d Developer
	if err := First(db.DB, &d); err != nil {
		t.Fatal(err)
	}

	if d.Name != dev.Name {
		t.Fatalf("Expected %q, got %q.\n%v", dev.Name, d.Name, nil)
	}
}

func TestDeveloperShouldReturnTenEntries(t *testing.T) {
	var db internal.Db = internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator DeveloperMigrator = DeveloperMigrator{}
	migrator.Clear(db.DB)
	migrator.Migrate(db.DB)

	var seeder DeveloperSeeder = DeveloperSeeder{}
	seeder.Seed(db.DB)

	var devs []*Developer = []*Developer{}

	if err := List(db.DB, 10, &devs); err != nil {
		t.Fatal(err)
	}

	if len(devs) != 10 {
		t.Fatalf("Expected length of %d, got %d.", 10, len(devs))
	}
}

func TestDeveloperShouldPaginate(t *testing.T) {
	var db internal.Db = internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator DeveloperMigrator = DeveloperMigrator{}
	migrator.Clear(db.DB)
	migrator.Migrate(db.DB)

	var seeder DeveloperSeeder = DeveloperSeeder{}
	seeder.Seed(db.DB)

	var devs []*Developer = []*Developer{}

	if err := Paginate(db.DB, 2, 5, &devs); err != nil {
		t.Fatal(err)
	}

	if len(devs) != 5 {
		t.Fatalf("Expected length of %d, got %d.", 5, len(devs))
	}

	if devs[0].ID != 6 {
		t.Fatalf("Expected ID of %d, got %d.", 6, devs[0].ID)
	}
}
