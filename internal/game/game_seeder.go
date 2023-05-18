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
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.mobygames.com%2Fimages%2Fcovers%2Fl%2F270570-final-fantasy-xiv-online-a-realm-reborn-windows-front-cover.jpg&f=1&nofb=1&ipt=1adf1510d2ea11c715192bef92b9474621918086d5306bbe7b95a3fd071a93ea&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Stardew Valley",
		Description: "A videogame about farming and slaying slimes.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse3.mm.bing.net%2Fth%3Fid%3DOIP.B8zc83uwXUA1pkDg27FfdQHaLH%26pid%3DApi&f=1&ipt=b47f8a68a3b1a29b4478088bdebf031c4d5c6e3d8628b5a6d64fdb5979674958&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Battlefield 2042",
		Description: "A game that could've been great but turned out mediocre.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.windowscentral.com%2Fsites%2Fwpcentral.com%2Ffiles%2Fstyles%2Fsmall%2Fpublic%2Ffield%2Fimage%2F2021%2F06%2Fbattlefield-2042-cover-art-standard.jpg&f=1&nofb=1&ipt=8eb44f51d5c114e7ed2c85113c994344ca3a44e5692cc11422105d8d4c7b3ccc&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Elden Ring",
		Description: "AAAAAUUUUGH THE EELDEEEEEN RIIIINGGGGG.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.mobygames.com%2Fimages%2Fcovers%2Fl%2F775869-elden-ring-xbox-one-front-cover.jpg&f=1&nofb=1&ipt=01743f407be933235d1cbcd4162ae066a40425f7439b05fd6e90799be336856a&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Yakuza 0",
		Description: "A grim and realistic videogame about the Japanese yakuza.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fcdn.wikimg.net%2Fen%2Fstrategywiki%2Fimages%2F8%2F89%2FYakuza_0_cover.jpg&f=1&nofb=1&ipt=901a23274a9ed0acb8c0f3e338a9a6d78c741c2b4b003e361ea94c4169e957ea&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Judgment",
		Description: "A videogame about serious private detectives.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse3.mm.bing.net%2Fth%3Fid%3DOIP.NV-s3kADi8iTel62Cqw-wAHaHa%26pid%3DApi&f=1&ipt=647432ac5745667d1c13312dac2363113ef6eed0e763c741451de70db6fae72e&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Lost Judgment",
		Description: "A videogame about dancing with bully victims.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fimage.api.playstation.com%2Fvulcan%2Fap%2Frnd%2F202106%2F1115%2FwQDzz8HqbaRvFqVKWDiy9qdQ.png&f=1&nofb=1&ipt=b20b0f8b027a1527fe557cc46e0e75d863feba180997c7576e400ea895114c12&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Minecraft",
		Description: "We are mining diamonds. O-wowowwowoohhohohoooo.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.mobygames.com%2Fimages%2Fcovers%2Fl%2F489736-minecraft-windows-apps-front-cover.jpg&f=1&nofb=1&ipt=45760016337307744cb9467dc2e95d6d2a34e5ea613a9f00f0cfffa625c8095b&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "Osu!",
		Description: "Click the circles. Till you die.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.redd.it%2Fbtji0qmnhtv31.png&f=1&nofb=1&ipt=838c72b6dc97b7ff7899dce6656dacb7b83bb2ca9aff22abba25b1aa7b98e305&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
	{
		Title:       "The Witcher 3: Wild Hunt",
		Description: "Ard Skellig truly is the best thing to ever exist.",
		Thumbnail:   "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.mobygames.com%2Fimages%2Fcovers%2Fl%2F392214-the-witcher-3-wild-hunt-xbox-one-front-cover.png&f=1&nofb=1&ipt=2c77ad1bdcba09aaadc58789d4a7c3cf4fc638e432ee59c25d791fe75b75254b&ipo=images",
		Developer: []*developer.Developer{
			{},
		},
	},
}

func (s GameSeeder) Seed(db *sql.DB) error {
	var query string = `
        INSERT INTO games (
            title, description, thumbnail
        ) VALUES (
            $1,
            $2,
            $3
        );
    `

	for _, g := range games {
		if _, err := db.Exec(query, g.Title, g.Description, g.Thumbnail); err != nil {
			return err
		}
	}

	return nil
}
