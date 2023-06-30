package seeders

import (
	"errors"
	"fmt"

	"github.com/bjvanbemmel/game-store/internal/database"
)

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
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Survival'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Online Co-Op'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Action RPG'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Cyberpunk'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Horror'
            );
        `, `
            INSERT INTO genres (
                name
            ) VALUES (
                'Sci-fi'
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
