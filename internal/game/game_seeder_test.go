package game

import (
	"database/sql"
	"testing"
)

func TestGameSeederSuccessfullySeeds(t *testing.T) {
	var seeder GameSeeder = GameSeeder{}

	err := seeder.Seed(sql.DB{})
	if err != nil {
		t.Fatal(err)
	}
}
