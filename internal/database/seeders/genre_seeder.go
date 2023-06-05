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
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'FPS'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'RPG'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'JRPG'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'MMORPG'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Racing'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Simulation'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Story Rich'
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
