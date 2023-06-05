package seeders

import "github.com/bjvanbemmel/game-store/internal/database"

type GameSeeder struct{}

func (s GameSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO games (
                title, price, thumbnail, description
            ) VALUES (
                'Final Fantasy XIV: A Realm Reborn',
                19.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.gamereleasedates.net%2Fimages%2Fcovers%2Fpc%2Fcover-pc-final-fantasy-xiv-a-realm-reborn.jpg&f=1&nofb=1&ipt=7ab31827c31c15e3c26a8715e05b1285a679f58a72f1d80c10b2de8346e1fa7b&ipo=images',
                (E'An MMORPG that will eat up your life. Y\'shtola is pure and utter waifu material.')
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description
            ) VALUES (
                'Yakuza 7: Like A Dragon',
                39.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmedia.senscritique.com%2Fmedia%2F000019514438%2Fsource_big%2FYakuza_Like_a_Dragon.png&f=1&nofb=1&ipt=b60654a54a1b4376ec07d81e805240dadcd9e9218ac7c13f801a3368a19fb7aa&ipo=images',
                'A totally grim and serious videogame about organized crime taking place in Japan.'
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description
            ) VALUES (
                'Assetto Corsa',
                19.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fcdkeyprices.com%2Fimages%2Fgames%2F5632433%2Fcover.jpg&f=1&nofb=1&ipt=3b65db0133b674285d182f8cefb975dbfdd5dcce6f5892380496f9d72ee6ce40&ipo=images',
                'Nothing feels better than driving a pimped Skyline through the Tokyo expressway while dodging traffic.'
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description
            ) VALUES (
                'NieR: Automata',
                39.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.gamekeys.se%2Fwp-content%2Fuploads%2F2021%2F01%2Fnier-automata-game-of-the-yorha-edition_cover_original.jpg&f=1&nofb=1&ipt=c22e63cd20e38426abab9285a5b4d5ac61e69c3446d8aca7ff28bac88a51dbbf&ipo=images',
                'Become. As. Gods. このままじゃ、だめ。'
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
