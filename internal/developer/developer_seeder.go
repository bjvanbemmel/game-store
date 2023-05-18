package developer

import "database/sql"

type DeveloperSeeder struct{}

var developers []Developer = []Developer{
	{
		Name: "Square Enix",
	},
	{
		Name: "Concerned Ape",
	},
	{
		Name: "Chucklefish",
	},
	{
		Name: "Electronic Arts",
	},
	{
		Name: "DICE",
	},
	{
		Name: "Bandai Namco",
	},
	{
		Name: "From Software",
	},
	{
		Name: "SEGA",
	},
	{
		Name: "Ryu Ga Gotoku Studio",
	},
	{
		Name: "Xbox Game Studios",
	},
	{
		Name: "Mojang Studios",
	},
	{
		Name: "Peppy",
	},
	{
		Name: "CD Projekt Red",
	},
}

func (s DeveloperSeeder) Seed(db *sql.DB) error {
	var query string = `
        INSERT INTO developers (
            name
        ) VALUES (
            $1
        )
    `

	for _, d := range developers {
		if _, err := db.Exec(query, d.Name); err != nil {
			return err
		}
	}

	return nil
}
