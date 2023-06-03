package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type DeveloperSeeder struct{}

func (s DeveloperSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO developers (
                name
            ) VALUES (
                'Square Enix'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Ryu Ga Gotoku'
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
