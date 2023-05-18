package cmd

import (
	"github.com/bjvanbemmel/game-store/internal/developer"
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

	var gameMigrator game.GameMigrator = game.GameMigrator{}
	if err := gameMigrator.Clear(db.DB); err != nil {
		return err
	}
	if err := gameMigrator.Migrate(db.DB); err != nil {
		return err
	}

	var devMigrator developer.DeveloperMigrator = developer.DeveloperMigrator{}
	if err := devMigrator.Clear(db.DB); err != nil {
		return err
	}
	if err := devMigrator.Migrate(db.DB); err != nil {
		return err
	}

	return nil
}
