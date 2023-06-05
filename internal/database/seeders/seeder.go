package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

var (
	db database.Db = database.Db{}
)

func Seed() error {
	var err error

	err = DeveloperSeeder{}.Seed()
	if err != nil {
		return err
	}

	err = CategorySeeder{}.Seed()
	if err != nil {
		return err
	}

	err = GameSeeder{}.Seed()
	if err != nil {
		return err
	}

	err = GameDeveloperSeeder{}.Seed()
	if err != nil {
		return err
	}

	err = GameCategorySeeder{}.Seed()
	if err != nil {
		return err
	}

	return nil
}
