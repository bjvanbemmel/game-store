package game

import (
	"database/sql"
	"testing"

	"github.com/bjvanbemmel/game-store/internal/developer"
)

func TestGameCanInsertAndFindResult(t *testing.T) {
	var dsn string = "file:unit_tests.db?cache=shared&mode=memory"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	var migrator GameMigrator = GameMigrator{}
	migrator.Migrate(db)

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

	err = Insert(db, &game)
	if err != nil {
		t.Fatal(err)
	}

	var g Game = Game{}
	if err := First(db, &g); err != nil {
		t.Fatal(err)
	}
}

func TestGameShouldReturnTenEntries(t *testing.T) {
	var dsn string = "file:unit_tests.db?cache=shared&mode=memory"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	var migrator GameMigrator = GameMigrator{}
	migrator.Migrate(db)

	var seeder GameSeeder = GameSeeder{}
	seeder.Seed(db)

	var games []*Game = []*Game{}

	err = List(db, 10, &games)
	if err != nil {
		t.Fatal(err)
	}

	if len(games) != 10 {
		t.Fatalf("Expected length of %d, got %d.\n%v", 10, len(games), err)
	}
}

func TestGameShouldPaginateFromFiveToTen(t *testing.T) {
	var dsn string = "file:unit_tests.db?cache=shared&mode=memory"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	var migrator GameMigrator = GameMigrator{}
	migrator.Migrate(db)

	var seeder GameSeeder = GameSeeder{}
	seeder.Seed(db)

	var games []*Game = []*Game{}

	err = Paginate(db, 2, 5, &games)
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
