package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type GameGenreSeeder struct{}

func (s GameGenreSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                1, 6
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                1, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                1, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                2, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                2, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                2, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                3, 7
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                3, 8
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                4, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                4, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                4, 9
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
