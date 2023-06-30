package seeders

import (
	"errors"
	"fmt"

	"github.com/bjvanbemmel/game-store/internal/database"
)

type GameSeeder struct{}

func (s GameSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'Final Fantasy XIV: A Realm Reborn',
                19.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fcdkeyprices.com%2Fimages%2Fgames%2F5634889%2Fcover.jpg&f=1&nofb=1&ipt=60404141980ef453c56f92cc2429bbd0362fb0c6ed8fa65464d3d3510a5a3c97&ipo=images',
                (E'An MMORPG that will eat up your life. Y\'shtola is pure and utter waifu material.'),
                '2013-08-27',
                95.23
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'Yakuza 7: Like A Dragon',
                39.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmedia.senscritique.com%2Fmedia%2F000019514438%2Fsource_big%2FYakuza_Like_a_Dragon.png&f=1&nofb=1&ipt=b60654a54a1b4376ec07d81e805240dadcd9e9218ac7c13f801a3368a19fb7aa&ipo=images',
                'A totally grim and serious videogame about organized crime taking place in Japan.',
                '2020-11-10',
                85.79
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'Assetto Corsa',
                19.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fcdkeyprices.com%2Fimages%2Fgames%2F5632433%2Fcover.jpg&f=1&nofb=1&ipt=3b65db0133b674285d182f8cefb975dbfdd5dcce6f5892380496f9d72ee6ce40&ipo=images',
                'Nothing feels better than driving a pimped Skyline through the Tokyo expressway while dodging traffic.',
                '2014-12-19',
                90.03
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'NieR: Automata',
                39.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.gamekeys.se%2Fwp-content%2Fuploads%2F2021%2F01%2Fnier-automata-game-of-the-yorha-edition_cover_original.jpg&f=1&nofb=1&ipt=c22e63cd20e38426abab9285a5b4d5ac61e69c3446d8aca7ff28bac88a51dbbf&ipo=images',
                'Become. As. Gods. このままじゃ、だめ。',
                '2017-03-17',
                98.30
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'The Witcher® 3: Wild Hunt',
                39.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.hMNiTiROyEPyi_TC3EpaqQHaLH%26pid%3DApi&f=1&ipt=f72d5478a7d52e5d979a41fe53ddf3511d025ffd8fe3caf3be1bbf211abfac8d&ipo=images',
                'You are Geralt of Rivia, mercenary monster slayer. Before you stands a war-torn, monster-infested continent you can explore at will. Your current contract? Tracking down Ciri — the Child of Prophecy, a living weapon that can alter the shape of the world.',
                '2015-05-18',
                93.91
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'NieR Replicant™ ver.1.22474487139...',
                59.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.gamerclick.it%2Fimmagini%2Fvideogame%2FNier_Replicant%2Fcover%2FNier_Replicant-cover.jpg&f=1&nofb=1&ipt=676350c2043997b35f55f66d0317c72ab9aafd6dabe6ab9ff20b0662d1c17a8d&ipo=images',
                'The upgraded prequel of NieR:Automata. A kind young man sets out with Grimoire Weiss, a strange talking book, to search for the "Sealed verses" in order to save his sister Yonah, who fell terminally ill to the Black Scrawl.',
                '2021-04-23',
                84.12
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'Valheim',
                19.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fpreview.redd.it%2Fcb6vhjj30ap61.gif%3Fformat%3Dpng8%26s%3D7fede5f0fb30bd3aa14dfdb16eb56e8e8412dedc&f=1&nofb=1&ipt=a961730d85250170fac8fe050c4480b51ff9c1541f713909dc1e67171ec68ea1&ipo=images',
                'A brutal exploration and survival game for 1-10 players, set in a procedurally-generated purgatory inspired by viking culture. Battle, build, and conquer your way to a saga worthy of Odin’s patronage! ',
                '2021-02-02',
                90.03
            );
        `, `
            INSERT INTO games (
                title, thumbnail, description, release_date, rating
            ) VALUES (
                'Lost Ark',
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fhowlongtobeat.com%2Fgames%2F26127_Lost_Ark.jpg&f=1&nofb=1&ipt=9796a7f8b803d7aa45cb539b6c5260afe8c4f4f03a1328feaa90ad134c72c3d8&ipo=images',
                'Embark on an odyssey for the Lost Ark in a vast, vibrant world: explore new lands, seek out lost treasures, and test yourself in thrilling action combat in this action-packed free-to-play RPG.',
                '2022-02-11',
                52.86
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'Cyberpunk 2077',
                59.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.mobygames.com%2Fimages%2Fcovers%2Fl%2F567922-cyberpunk-2077-xbox-one-front-cover.jpg&f=1&nofb=1&ipt=a3f68d17f7277e005c4dd8616d0cc68e777345002edfff10690acf0e6eda6dcf&ipo=images',
                'Cyberpunk 2077 is an open-world, action-adventure RPG set in the dark future of Night City — a dangerous megalopolis obsessed with power, glamor, and ceaseless body modification.',
                '2020-12-10',
                79.52
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'Ghostwire: Tokyo',
                59.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2F1.bp.blogspot.com%2F-W2Z-i7l0ayY%2FYCBtNdyBlzI%2FAAAAAAAADBY%2FbvbOXgzB2PEK1bUuZfCRnqulU0XaFFcbwCLcBGAsYHQ%2Fs1867%2Fghostwire-tokyo-cover.jpg&f=1&nofb=1&ipt=39c44fe3e143afa4c51c08164a700865dfbae5395b73f3a782a74fd0f52692ba&ipo=images',
                (E'Tokyo\'s population has vanished, and deadly supernatural forces prowl the streets. Use an arsenal of elemental abilities to unravel the truth behind the disappearance and save Tokyo.'),
                '2022-03-25',
                75.18
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'Prey',
                29.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fpostergg.savebutonu.com%2Fwp-content%2Fuploads%2F2019%2F12%2FPrey-1-1448x2048.jpg&f=1&nofb=1&ipt=c37a32aed4159984a29a3d05672c0bc947e2b51cc0b8de9e4051dc52f07505df&ipo=images',
                'In Prey, you awaken aboard Talos I, a space station orbiting the moon in the year 2032. You are the key subject of an experiment meant to alter humanity forever – but things have gone terribly wrong. The space station has been overrun by hostile aliens and you are now being hunted.',
                '2017-05-4',
                89.45
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'ELDEN RING',
                59.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.mobygames.com%2Fimages%2Fcovers%2Fl%2F775869-elden-ring-xbox-one-front-cover.jpg&f=1&nofb=1&ipt=ddeb56331884f6743c06adbbf892ed7eab81fa0b3b53a978d7069c514528bce2&ipo=images',
                'THE NEW FANTASY ACTION RPG. Rise, Tarnished, and be guided by grace to brandish the power of the Elden Ring and become an Elden Lord in the Lands Between.',
                '2022-02-25',
                96.04
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'SIGNALIS',
                19.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmediamaster.vandal.net%2Fm%2F122330%2Fsignalis-2022101810373354_1.jpg&f=1&nofb=1&ipt=cd1edf5bb2e5126cf1b2286dc4e9ecf9681a967c7965a0e2e0bd3cda7eb3c3ed&ipo=images',
                'A classic survival horror experience set in a dystopian future where humanity has uncovered a dark secret. Unravel a cosmic mystery, escape terrifying creatures, and scavenge an off-world government facility as Elster, a technician Replika searching for her lost dreams.',
                '2022-10-27',
                95.55
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date, rating
            ) VALUES (
                'CODE VEIN',
                59.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.mobygames.com%2Fimages%2Fcovers%2Fl%2F594953-code-vein-xbox-one-front-cover.jpg&f=1&nofb=1&ipt=e7b27bf606a067c073094e4f891a0a6e038a747790ef51b2a559ef04050f6037&ipo=images',
                'In the face of certain death, we rise. Team up and embark on a journey to the ends of hell to unlock your past and escape your living nightmare in CODE VEIN.',
                '2019-09-27',
                72.98
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date
            ) VALUES (
                'ARMORED CORE™ VI FIRES OF RUBICON™',
                59.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fstatic.actugaming.net%2Fmedia%2F2022%2F12%2Farmored-core-6-fires-of-rubicon-jaquette.jpg&f=1&nofb=1&ipt=447b999bf5787e053405024fd10782d85135b711987e8e40906508b16eed706b&ipo=images',
                (E'A new action game based on the concept of the ARMORED CORE series that uses the knowledge gained from FromSoftware\'s recent action game development.'),
                '2023-08-25'
            );
        `, `
            INSERT INTO games (
                title, price, thumbnail, description, release_date
            ) VALUES (
                'STARFIELD',
                69.99,
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fstatic.actugaming.net%2Fmedia%2F2018%2F06%2Fstarfield-game-cover-e1623683439466.jpg&f=1&nofb=1&ipt=7713acbad095fc407842873e79b4052c71e470153b551141783f17b73d570523&ipo=images',
                'Starfield is the first new universe in 25 years from Bethesda Game Studios, the award-winning creators of The Elder Scrolls V: Skyrim and Fallout 4.',
                '2023-09-06'
            );
        `, `
            INSERT INTO games (
                title, thumbnail, description, release_date
            ) VALUES (
                'BLUE PROTOCOL',
                'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fxn--80aigig6c.xn--p1acf%2Fwp-content%2Fuploads%2F2020%2F09%2Fblue-protocol-cover.jpg&f=1&nofb=1&ipt=9887a700ca574c29e74d783c71951e239e417ef1e6b6f5a7c3f8b28447e69027&ipo=images',
                'Featuring a dynamic action combat system, deep character customization, and an epic story set in a vast world, BLUE PROTOCOL is an Online Action RPG where you become the hero of your very own anime adventure. ',
                '2024-01-01'
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
