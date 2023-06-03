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

	err = AddCategoriesTable{}.Migrate()
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

	err = AddGameCategoryPivotTable{}.Migrate()
	if err != nil {
		return err
	}

	return nil
}
