package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type GameDeveloperSeeder struct{}

func (s GameDeveloperSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                1, 1
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                1, 2
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
