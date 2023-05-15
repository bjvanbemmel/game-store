package game

import "database/sql"

type GameSeeder struct{}

func (s GameSeeder) Seed(db sql.DB) error {
	return nil
}
