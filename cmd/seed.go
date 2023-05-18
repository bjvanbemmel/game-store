package cmd

import (
	"github.com/bjvanbemmel/game-store/internal/developer"
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

	var gameSeeder game.GameSeeder = game.GameSeeder{}
	if err := gameSeeder.Seed(db.DB); err != nil {
		return err
	}

	var devSeeder developer.DeveloperSeeder = developer.DeveloperSeeder{}
	if err := devSeeder.Seed(db.DB); err != nil {
		return err
	}

	return nil
}
