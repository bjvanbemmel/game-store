package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type GameSeeder struct{}

func (s GameSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO games (
                title, price
            ) VALUES (
                'Final Fantasy XIV: A Realm Reborn',
                19.99
            );
        `, `
            INSERT INTO games (
                title, price
            ) VALUES (
                'Yakuza 7: Like A Dragon',
                39.99
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
