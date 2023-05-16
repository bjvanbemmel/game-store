package game

import (
	"testing"

	"github.com/bjvanbemmel/game-store/internal"
	"github.com/bjvanbemmel/game-store/internal/developer"
)

func TestGameCanInsertAndFindResult(t *testing.T) {
	var db internal.Db = internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}
	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator GameMigrator = GameMigrator{}
	migrator.Migrate(db.DB)

	var game Game = Game{
		Title:       "Stardew Valley",
		Description: "A videogame",
		Thumbnail:   "https://duckduckgo.com/",
		Developer: []*developer.Developer{
			{
				Name: "Concerned Ape",
			},
		},
	}

	t.Log(db.ConnectionString)

	err := Insert(db.DB, &game)
	if err != nil {
		t.Fatal(err)
	}

	var g Game = Game{}
	if err := First(db.DB, &g); err != nil {
		t.Fatal(err)
	}
}

func TestGameShouldReturnTenEntries(t *testing.T) {
	db := internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}

	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator GameMigrator = GameMigrator{}
	migrator.Clear(db.DB)
	migrator.Migrate(db.DB)

	var seeder GameSeeder = GameSeeder{}
	seeder.Seed(db.DB)

	var games []*Game = []*Game{}

	err := List(db.DB, 10, &games)
	if err != nil {
		t.Fatal(err)
	}

	if len(games) != 10 {
		t.Fatalf("Expected length of %d, got %d.\n%v", 10, len(games), err)
	}
}

func TestGameShouldPaginateFromFiveToTen(t *testing.T) {
	db := internal.Db{}
	if err := db.Environment("test"); err != nil {
		t.Fatal(err)
	}

	if err := db.Connect(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var migrator GameMigrator = GameMigrator{}
	migrator.Clear(db.DB)
	migrator.Migrate(db.DB)

	var seeder GameSeeder = GameSeeder{}
	seeder.Seed(db.DB)

	var games []*Game = []*Game{}

	err := Paginate(db.DB, 2, 5, &games)
	if err != nil {
		t.Fatal(err)
	}

	if len(games) != 5 {
		t.Fatalf("Expected length of %d, got %d.\n%v", 10, len(games), err)
	}

	if games[0].ID != 6 {
		t.Fatalf("Expected ID of %d, got %d.\n%v", 4, games[0].ID, err)
	}
}
