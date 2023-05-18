package developer

import "database/sql"

type DeveloperMigrator struct{}

func (m DeveloperMigrator) Migrate(db *sql.DB) error {
	var query string = `
        CREATE TABLE IF NOT EXISTS developers(
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL
        )
    `

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (m DeveloperMigrator) Clear(db *sql.DB) error {
	var query string = `
        DROP TABLE IF EXISTS developers;
    `

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
