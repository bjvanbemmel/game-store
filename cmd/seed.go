package cmd

import (
	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/bjvanbemmel/game-store/internal/database/seeders"
)

type Seed struct{}

func (s Seed) Execute() error {
	if err := database.Context.Connect(); err != nil {
		return err
	}

	if err := seeders.Seed(); err != nil {
		return err
	}

	return nil
}
