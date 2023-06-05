package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type CategorySeeder struct{}

func (s CategorySeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO categories (
                name
            ) VALUES (
                'Action'
            );
        `, `
            INSERT INTO categories (
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
