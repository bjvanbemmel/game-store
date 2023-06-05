package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type GenreSeeder struct{}

func (s GenreSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO genres (
                name
            ) VALUES (
                'Action'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Adventure'
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
