package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type GameCategorySeeder struct{}

func (s GameCategorySeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO game_category (
                game_id, category_id
            ) VALUES (
                1, 1
            );
        `, `
            INSERT INTO game_category (
                game_id, category_id
            ) VALUES (
                1, 1
            );
        `,
	}

	for _, q := range queries {
		if _, err := database.Context.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
