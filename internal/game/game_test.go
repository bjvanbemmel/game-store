package game

import (
	"database/sql"
	"testing"
)

func TestGameInsertsProperly(t *testing.T) {
	var game Game = Game{
		Title: "Test",
	}

	err := Insert(sql.DB{}, &game)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGameFirstReturnsValidRecord(t *testing.T) {
	var game Game = Game{}

	err := First(sql.DB{}, &game)
	if err != nil {
		t.Fatal(err)
	}

	if game.Title != "Test" {
		t.Fatalf("Expected %q, got %q.\n%v", "Test", game.Title, err)
	}
}

func TestGameShouldReturnTenEntries(t *testing.T) {
	var games []*Game = []*Game{}

	err := List(sql.DB{}, 10, games)
	if err != nil {
		t.Fatal(err)
	}

	if len(games) != 10 {
		t.Fatalf("Expected length of %q, got %q.\n%v", 10, len(games), err)
	}
}

func TestGameShouldPaginateFromTenToTwenty(t *testing.T) {
	var games []*Game = []*Game{}

	err := Paginate(sql.DB{}, 10, 10, games)
	if err != nil {
		t.Fatal(err)
	}

	if len(games) != 10 {
		t.Fatalf("Expected length of %q, got %q.\n%v", 10, len(games), err)
	}

	if games[0].ID != 11 {
		t.Fatalf("Expected ID of %q, got %q.\n%v", 11, games[0].ID, err)
	}
}
