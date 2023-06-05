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
                2, 2
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                2, 3
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                3, 4
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                4, 5
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                4, 1
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
