package game

import "database/sql"

type GameMigrator struct{}

func (m GameMigrator) Migrate(db *sql.DB) error {
	var query string = `
        CREATE TABLE IF NOT EXISTS games(
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            description VARCHAR(255) NOT NULL,
            thumbnail VARCHAR(255) NOT NULL
        );
    `

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (m GameMigrator) Clear(db *sql.DB) error {
	var query string = `
        DROP TABLE IF EXISTS games;
        `

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
