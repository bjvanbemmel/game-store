package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type GameGenreSeeder struct{}

func (s GameGenreSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                1, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
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
