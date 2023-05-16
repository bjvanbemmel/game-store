package cmd

import (
	"github.com/bjvanbemmel/game-store/internal/game"
)

type Seed struct{}

func (s Seed) Execute() error {
	if err := db.Environment("prod"); err != nil {
		return err
	}

	if err := db.Connect(); err != nil {
		return err
	}

	var seeder game.GameSeeder = game.GameSeeder{}
	err := seeder.Seed(db.DB)

	return err
}
