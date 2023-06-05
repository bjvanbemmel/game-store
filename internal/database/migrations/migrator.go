package migrations

import "github.com/bjvanbemmel/game-store/internal/database"

var (
	db database.Db = database.Db{}
)

func Migrate() error {
	var err error

	err = AddDevelopersTable{}.Migrate()
	if err != nil {
		return err
	}

	err = AddGenreTable{}.Migrate()
	if err != nil {
		return err
	}

	err = AddGamesTable{}.Migrate()
	if err != nil {
		return err
	}

	err = AddGameDeveloperPivotTable{}.Migrate()
	if err != nil {
		return err
	}

	err = AddGameGenrePivotTable{}.Migrate()
	if err != nil {
		return err
	}

	return nil
}
