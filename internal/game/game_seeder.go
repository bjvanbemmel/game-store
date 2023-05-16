package game

import (
	"database/sql"

	"github.com/bjvanbemmel/game-store/internal/developer"
)

type GameSeeder struct{}

var games []Game = []Game{
	{
		Title:       "Final Fantasy XIV: A Realm Reborn",
		Description: "A videogame about heroic weaboos.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Stardew Valley",
		Description: "A videogame about farming and slaying slimes.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Battlefield 2042",
		Description: "A game that could've been great but turned out mediocre.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Elden Ring",
		Description: "AAAAAUUUUGH THE EELDEEEEEN RIIIINGGGGG.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Yakuza 0",
		Description: "A grim and realistic videogame about the Japanese yakuza.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Judgment",
		Description: "A videogame about serious private detectives.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Lost Judgment",
		Description: "A videogame about dancing with bully victims.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Minecraft",
		Description: "We are mining diamonds. O-wowowwowoohhohohoooo.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Osu!",
		Description: "Click the circles. Till you die.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "The Witcher 3: Wild Hunt",
		Description: "Ard Skellig truly is the best thing to ever exist.",
		Thumbnail:   "https://placehold.er",
		Developer: []*developer.Developer{
			{},
		},
	},
}

func (s GameSeeder) Seed(db *sql.DB) error {
	var query string = `
        INSERT INTO games (
            id, title, description, thumbnail
        ) VALUES (
            NULL,
            ?,
            ?,
            ?
        );
    `

	for _, g := range games {
		if _, err := db.Exec(query, g.Title, g.Description, g.Thumbnail); err != nil {
			return err
		}
	}

	return nil
}
