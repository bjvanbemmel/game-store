package cmd

import (
	"github.com/bjvanbemmel/game-store/internal/database"
	"github.com/bjvanbemmel/game-store/internal/database/migrations"
)

type Migrate struct{}

func (m Migrate) Execute() error {
	if err := database.Context.Connect(); err != nil {
		return err
	}

	if err := migrations.Migrate(); err != nil {
		return err
	}

	return nil
}
