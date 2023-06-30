package seeders

import (
	"errors"
	"fmt"

	"github.com/bjvanbemmel/game-store/internal/database"
)

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
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                5, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                5, 4
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                5, 9
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                6, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                6, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                6, 9
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                7, 10
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                7, 11
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                8, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                8, 6
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                8, 12
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                9, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                9, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                9, 13
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                9, 9
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                9, 15
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                10, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                10, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                10, 14
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                11, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                11, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                11, 15
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                12, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                12, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                12, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                13, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                13, 14
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                13, 15
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                14, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                14, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                15, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                15, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                15, 15
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                16, 4
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                16, 9
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                16, 15
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                17, 6
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                17, 5
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                17, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                17, 2
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                17, 1
            );
        `, `
            INSERT INTO game_genre (
                game_id, genre_id
            ) VALUES (
                17, 4
            );
        `,
	}

	for _, q := range queries {
		if _, err := database.Context.Exec(q); err != nil {
			return errors.New(fmt.Sprintln(err.Error(), q))
		}
	}

	return nil
}
