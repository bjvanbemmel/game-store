package seeders

import (
	"errors"
	"fmt"

	"github.com/bjvanbemmel/game-store/internal/database"
)

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
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Sega'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Kunos Simulazioni'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Platinum Games'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'CD PROJEKT RED'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Toylogic Inc.'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Iron Gate AB'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Coffee Stain Publishing'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Smilegate RPG'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Amazon Games'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Bethesda Softworks'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Tango Gameworks'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Arkane Studios'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'FromSoftware Inc.'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Bandai Namco Entertainment'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'rose-engine'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'Humble Games'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'PLAYISM'
            );
        `, `
            INSERT INTO developers (
                name
            ) VALUES (
                'NEOWIZ'
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
