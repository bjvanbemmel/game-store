package cmd

import (
	"github.com/bjvanbemmel/game-store/internal/game"
)

type Migrate struct{}

func (m Migrate) Execute() error {
	if err := db.Environment("prod"); err != nil {
		return err
	}

	if err := db.Connect(); err != nil {
		return err
	}

	var migrator game.GameMigrator = game.GameMigrator{}
	if err := migrator.Clear(db.DB); err != nil {
		return err
	}

	err := migrator.Migrate(db.DB)

	return err
}
