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
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                5, 6
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                6, 1
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                6, 7
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                7, 8
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                7, 9
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                8, 10
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                8, 11
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                9, 6
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                10, 12
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                10, 13
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                11, 12
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                11, 14
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                12, 15
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                12, 16
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                13, 17
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                13, 18
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                13, 19
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                14, 16
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                15, 15
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                15, 16
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                16, 12
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                17, 16
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                17, 11
            );
        `, `
            INSERT INTO game_developer (
                game_id, developer_id
            ) VALUES (
                18, 20
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
