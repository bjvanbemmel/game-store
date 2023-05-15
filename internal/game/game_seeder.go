package game

import "database/sql"

type GameSeeder struct{}

func (s GameSeeder) Seed(db *sql.DB) error {
	var query string = `
        INSERT INTO games (
            id, title, description, thumbnail
        ) VALUES (
            NULL,
            "Final Fantasy XIV: A Realm Reborn",
            "A videogame",
            "https://gamer.dev/"
        );
    `

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
