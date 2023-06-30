package seeders

import (
	"errors"
	"fmt"

	"github.com/bjvanbemmel/game-store/internal/database"
)

type MediaSeeder struct{}

func (s MediaSeeder) Seed() error {
	var queries []string = []string{`
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                2,
                'https://youtube.com/embed/h542YbZuwkQ'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_ca57d5322b116ffaab2de37958df59eba8c7807a.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_f1bc7a959795c1fad6bfc602ccbd0e311075a454.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_ddb71dc0c9a667de5c23e0b3823dad3590d5e9bb.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_912ffecab72b0ec4132e100582a6592e559ac3aa.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_d37db153382217ad479679691899588e8d173937.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_b13bbf3e5c841c759483c63604b0359080700dbd.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_ed8445bec9d85b79079f7ec46be38adc4068e3c5.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_e4686dc9cbe1b0d9d4d96bc029c766b195325209.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_634fd54775211fc96a911857f4baac7109985b5b.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                1,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/39210/ss_8d189c2d2e92705f09375ab1bc951196c4bc33c1.1920x1080.jpg?t=1684850999'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                2,
                'https://youtube.com/embed/hhqbNaxO82g'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_3672df00523861cd37b0f969d80604003ba14fd4.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_9cecfb713527a480f607bbde54c01763b18bf354.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_d8bcb8c72368ec09506d3a60d42ff2a1901e39f7.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_23d607ec8ceaca26edc88b6445d92ddcd111fea4.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_66151ac4259635b4f90e7945706498dacf6e98c1.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_b7336fc083ba36c6073c25659fe23075234b179a.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_8562245f8ca68976ffb5404023d148fb0cbc13e9.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_fe349bb52fa4a8243ba8d9b861bb52c3a4f4b5b0.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_a46dfebb4726aa042c76636131f8bbaa1a079fb9.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_5a17b24cc9582e458a3db0f5f9ecc3f956ded072.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                2,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1235140/ss_d3fe630add22247a497f0c2d22e568df66a39f48.1920x1080.jpg?t=1686319204'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                2,
                'https://youtube.com/embed/JfJNvlFa9TE'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_cc532907a14b12117825c609b9c941768ce87e51.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_b97afb7a10ad6455c01d0e2a191e105007cf7504.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_131963607a54afaf64688397d9742af37149148d.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_276544bf9b71fe701254f1d065e97c0be74f65e2.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_6c771fd2bff6a09d39541d0aee04a7f663927e0f.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_a098061e7aaf52410030a012d07582b11ff4cb72.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_61a962d326288c0dcd88edcd062a62b89e41d77d.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_c5b3513683139e9f97bff528e275ca3f89bd651e.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_5b7457e7ec2e04517be4e3e22c3eb1c182cd4b06.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_726fe3c27fd774eb790382c1b20066db139b61b9.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_a35de8332bc6aa22e637523e535bbbbc930ee390.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                3,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/244210/ss_4beda281a1dd82ccd82e32052dafdfc403c2916b.1920x1080.jpg?t=1669152181'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                2,
                'https://youtube.com/embed/wJxNhJ8fjFk'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_d0314b4c134329a483b5e43af951f60274abc66b.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_8b29f7e1ce9a8b9313dc3eb50dbe76a4cf94eef9.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_2c265df38c3d2d393d74ee8e74d79bdafa16b143.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_831e0e7c9d514393b711e9ed1d6796042521a80c.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_a6d164452c1aa00a0d7b7ca31aa76d787853b39e.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_edcb7633ff42d7200bcb240a1ebb1116d602d9fe.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_e926e3b5d440b4f244525745c7100edc2d717b85.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_02d2f3f2b7b7add8e6ad50d6b9325d05fa1d7bc7.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_b55c67ac11781513183391c18ea86819e047577d.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                4,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/524220/ss_c538e630c5cc224124104cc42ec6220ab90b5852.1920x1080.jpg?t=1682534695'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                2,
                'https://youtube.com/embed/c0i88t0Kacs'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_5710298af2318afd9aa72449ef29ac4a2ef64d8e.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_0901e64e9d4b8ebaea8348c194e7a3644d2d832d.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_112b1e176c1bd271d8a565eacb6feaf90f240bb2.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_d1b73b18cbcd5e9e412c7a1dead3c5cd7303d2ad.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_107600c1337accc09104f7a8aa7f275f23cad096.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_64eb760f9a2b67f6731a71cce3a8fb684b9af267.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_eda99e7f705a113d04ab2a7a36068f3e7b343d17.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_d5b80eb63c12a6484f26796f3e34410651bba068.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_b74d60ee215337d765e4d20c8ca6710ae2362cc2.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_dc55eeb409d6e187456a8e159018e8da098fa468.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_849ec8dcc6f8df1c0b2c509584c9fc9e51f88cfa.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_90a33d7764a2d23306091bfcb52265c3506b4fdb.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_07812c174bb6b96c29895ddc27404143df7aba6d.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_ed23139c916fdb9f6dd23b2a6a01d0fbd2dd1a4f.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_dc33eb233555c13fce939ccaac667bc54e3c4a27.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_908485cbb1401b1ebf42e3d21a860ddc53517b08.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_608af6cfe0eab3f37265550b391732a3e88d1a4f.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                5,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/292030/ss_8ac1dc847388e59b1be1c5ea5ca592d5861756cc.1920x1080.jpg?t=1675178392'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                2,
                'https://youtube.com/embed/sPcf4pfTqfY'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_2ff3ddb26c30b8397bce45abd0b4d847c3457259.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_f33bdbe53fa2d75e429f9b35a1299109c9ab8991.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_b7867dae1f1fa62a2cc82165c8c79eb6821782d6.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_8d7a5e5a0a7fe7782bf238763a2e29f8f6419e84.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_f7689ef6964eee2407d9e996bc73b380e3e7a56a.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_6740ec558094fafc86f0933264e50e796c21e9cf.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_eb6b583db6d9b3051c131ba748c768b2d6226277.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_c93af51b782d72cf081d4e8451836ad1716f63be.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_56ecf25d6d95441a23f5481f895771c2dcb9ac18.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_a2c004272402bf76b70ed13c920fbd3c85d43d94.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_0621ab7f10a0b23f4124509a13a8ace0487d856a.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_ada29940be1bdb0e8b70e64788044f4c2592657f.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_1c972f496956e106c3b9023c68ff309fd649407e.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_344cfa42952beb6d51540d43bdfcbeaa0949c9a7.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_2a5576351e2f84518dec11a796fbff609f4e5b9f.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_071609b98112f7621067dd5275ce0ead665f13f2.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_b3a1dfdbb9d4a47b6407de5c8b1dc75ddfcea931.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                6,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1113560/ss_942e62e6f1bf312148a7480f943babe2325a3714.1920x1080.jpg?t=1673963725'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                2,
                'https://youtube.com/embed/5mHRJ1KFe20'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_a600a7d4ca954543e22f571a9629521a13f82143.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_cd0262c5abf8a90ee5e1059acafc5a92b6be0e73.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_3db385fc1223914dadb199ac8682683a8c59454e.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_6032c9336c6cf9c1ebc914cedf022b38e97fd271.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_f6cbfeea728d557fa9f483685fea3205f08f5d9e.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_551f06a43b72609d7ca3cd63e93c58e949d58384.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_f4ea4f0f93cc8b38042f6d5916413da185ec221c.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_553301e2432883e9f026fe1dd0e91d7a8886d6f1.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_9b4abc60696de192c40064364a1395ad5074e5c3.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_15a86a1c4175f9392127265735c177a91535de65.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_ad81cc7ced8585ad344ba50e9d0b4bf9c597e62e.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_4ba939a383dc363c14e6cb7995322f39a5a26b07.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_988d92f03856895ef73a636ad20c2951086b865c.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_8db2e5f40b64e105c7174b3d3679fce6d7a04d20.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_b3c171e257dd1a3173254b1b61a7995dec204df4.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                7,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/892970/ss_64ae63dcfde6612266cf72884132fa144908b777.1920x1080.jpg?t=1676365340'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                2,
                'https://youtube.com/embed/Dhc30WRb0n0'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_f86c68e6fe904392d8a08231121f860814125f62.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_489519a55f1fcef124413a36232598017429a8d2.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_fd839dd6657095e6e61c45bcaea0fabbf1434096.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_c4789527b9087517504af18aeec0e2df769a54db.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_59b4b089b53a3459797ce9b0f21d6e4ff85f15c6.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_10fb9cd7c9be78e1ca7822a433003f28d8a14926.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_4e7802861ea7a1cb72e75a2003bf8769d47f6c4a.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_0835a9016d0c453675994ec218a3dd70a73529ea.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_b9c562ad66414ec37cf448f2f9d19d70009129ef.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                8,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1599340/ss_ad65376ee41e614c8e28993168c294cb600e4777.1920x1080.jpg?t=1678222947'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                2,
                'https://youtube.com/embed/UnA7tepsc7s'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_b529b0abc43f55fc23fe8058eddb6e37c9629a6a.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_8640d9db74f7cad714f6ecfb0e1aceaa3f887e58.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_9284d1c5b248726760233a933dbb83757d7d5d95.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_4bda6f67580d94832ed2d5814e41ebe018ba1d9e.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_e5a94665dbfa5a30931cff2f45cdc0ebea9fcebb.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_429db1d013a0366417d650d84f1eff02d1a18c2d.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_872822c5e50dc71f345416098d29fc3ae5cd26c1.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_ae4465fa8a44dd330dbeb7992ba196c2f32cabb1.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_f79fda81e6f3a37e0978054102102d71840f8b57.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_bb1a60b8e5061caef7208369f42c5c9d574c9ac4.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_a0c4e4015b5cf766d19bf97eee8b086183510e04.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_b20689e73e3ac19bcc5fad2c18d0353c769de144.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_ff3d920e254d18aa2a25d3765ac2ebe845efd208.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_0002f18563d313bdd1d82c725d411408ebf762b0.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_526123764d1c628caa1eb62c596f1b732f416c8c.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_284ba40590de8f604ae693631c751a0aefdc452e.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                9,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1091500/ss_9beef14102f164fa1163536d0fb3a51d0a2e4a3f.1920x1080.jpg?t=1680026109'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                2,
                'https://youtube.com/embed/2EDplmRKPLk'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_67c79411e5b14f370cefce560e84553027d210db.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_b62b51519d10a281d2d97a636e6952005b65fa32.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_eb2642ee902c4a34db5489fec084d1ee6db78105.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_c360bb24b6eee9ab0b313b512530c32ebebc19b5.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_72c5a29ebef19e5f7c046448e23c7a2a10d9d8f5.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_184714a58ea7d64df0d00efc8b60961b4d938289.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_b59b7c16cb92dece04459303befb1f256fe928ea.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_ead0ef8ed714f359508a48749d5a4bfa24df0dba.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_00e56cc260e1f6ab2d672eb24bc53b1228d99986.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_f7af029967284aab94a1a2b39fbcab87fe965dfa.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_8d6782c56f92bc48859800d59892d48c1cafa353.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_932734fa052e95e00aa5f7581d7824249553fccd.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_136db4c2671c60878815adbe1417d94cc7ab9835.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_2376b5ea9292666b2e9b08cf4111f1b4e40d1260.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                10,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1475810/ss_db65ea9601d4b855c013e713c4de078fca61d0f7.1920x1080.jpg?t=1681316334'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                2,
                'https://youtube.com/embed/LNHZ9WAertc'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_06b6c27c834b5639c54d470b3b5c711cf72a94af.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_ff6b4efb3add6ea9a1d67f5c6c0fae6661ed9fd8.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_976f1d181de3dc8bc86c78fe900c98be457d0942.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_788d370761d6b3c488f4314ea2752170621573c6.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_570af3ccb6cea896e3fb063d19ff87ce652c040b.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_8fcc1a7b13352ad4d0b13d6b32bcd77929f0bc2b.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_8f6126fa31c3cb52a77de2ff028413854e6189f6.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_e79a3faffdc0c289c6e1efc87ce3a94212b93e33.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_d657500c07a7ed972c891da21adb20aa05520a79.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                11,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/480490/ss_b0e3ef3f06f43ae02e70ac34ddb4b8dd693b1014.1920x1080.jpg?t=1594910513'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                2,
                'https://youtube.com/embed/qqiC88f9ogU'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_ae44317e3bd07b7690b4d62cc5d0d1df30367a91.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_e80a907c2c43337e53316c71555c3c3035a1343e.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_b70e156adf9e40aed24c10fb352b7813586e7290.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_3aec1455923ef49f4e777c2a94dbcd0256f77eb0.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_b6c4cdb36cebdbd52b97ab6e0851b7d3e41f03b3.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_e87a3e84890ab19f8995566e62762d5f8ed39315.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_1e3dfe515c04f4071207f01d62b85a1d6b560ced.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_3e556415d1bda00d749b2166ced264bec76f06ee.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_c372274833ae6e5437b952fa1979430546a43ad9.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_b87601dee58f4dbc36e40a8d803dc6a53ceefe07.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_8b58d96262fb0d62a482621b86c6ff85f4f57997.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_1011610a0e330c41a75ffd0b3a9a1bac3205c46a.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_41e2e8f3b0ad631e929e0c2ec3d1f21de883e98c.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_abd681cde3402155a35edb82919b7602cc7ec338.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_0b6e59057b02392b21dde62b4dde65d447e1fa3c.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_7523a8fc7775ae65cabd94d092ebecbd963258b6.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_ebb332e63dfab864c3bf3c3b1001b69f4da25f9f.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                12,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1245620/ss_75f25974c20b8704fe5ee246f74896b550088d3e.1920x1080.jpg?t=1683618443'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                2,
                'https://youtube.com/embed/4KFiOp2o4L8'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_a2603694154878b8260c1dd498a06168cad012a4.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_9d602346170b19121e4baec94f5fab54cc43637c.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_dc87789ecfa2f83dd58ac778866fbf7fdc1ec99a.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_bb676746c44db01c2bd7ca78616032984f958106.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_4ec0c13288054a99bc3987680a914ccda834e7f6.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_f7eb1a1944c4c9d3142bedb800ad3a43a6f82a60.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_947db45ff0fb6da0e2d9031f0f71a32b740ed612.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_dcf8d44a07289d9b8ca9d1e9b9ae2a002dd76f6e.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_c7fcc30c5e2cd2ddf44bc01b2d43964abba72076.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_5338061a9fa31755789e0a7698fadc2d4ab29b94.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                13,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/1262350/ss_af2cb31dc253d739c3ce2f6930be5d339e199be2.1920x1080.jpg?t=1684085470'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                2,
                'https://youtube.com/embed/tXvKDyuObbQ'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_8872f9fd79ef968f3a3023f17811c9f9133f7d64.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_2c00b08bf8a854b43757198e178be27ac46f50aa.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_4f2662706d87964e979c223f0264622574f151bc.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_6934095e7aae433f2e373fdf2319db7df1a899fb.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_d7671321de57658244e4bdc75c23dc63d78d0d7a.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_b5403477a94de6bb33663b1e4d9148be1f1a3b1d.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_2d1b27a5cb49e9ef1a89fd651e850174017d2dd6.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_ada0fd6894bcd9e1e316073f1e3ab61326d87b70.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_2b26bc9837a0b225aa7a7573c7800a8f91a6d510.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                14,
                1,
                'https://cdn.akamai.steamstatic.com/steam/apps/678960/ss_dd91c7aa4e9352aab3573ff57e48833704c8d05e.1920x1080.jpg?t=1676860712'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                2,
                'https://youtube.com/embed/SlSfr6Wa5sc'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_549f55589c10866bc31243d277324e31ad155b29.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_dcd98899647b45246cfb296aa5a3b40b2ae87e9e.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_c92110c39a9c64376af5e8da31a0e6ffa9747334.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_6648f7e39998ae173b2271c6a325d4295e6db785.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_f441df5b6d02d0cbc2635f29ec502d93058e025c.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_88767c4785d5b7e839a5434e988a8aee1836e1ad.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_15c13c66ab68af7fc733afb0bd0bd3358e6e842f.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_ef41adedb3c5c5977b60ae255440c4bdcc9b6f52.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_3302cd480fbba54627e226877d1d99f3ea323ff5.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_c8e974d1e0460034c4dce3b96605b3e2a6ed29db.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                15,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1888160/ss_560147479a040f97fcb507733f8fb85499d13e13.1920x1080.jpg?t=1687177060'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                2,
                'https://youtube.com/embed/kfYEiTdsyas'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_4887dc140a637684ddcfca518458668409f946dc.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_b2821283cb140cd5a6289a8160016b6a60d8f96e.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_68f15d580bf91971f637be5e464bc803482d78f7.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_aae99c177004bb5ec653d2fcb65a5d30489ec7b8.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_c8594798fadfd8e042b2fc8afff7bcf4872c5198.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_2288919a390c0147b7d2226354a61452016fd087.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_45c1dc3cd5399eb16230ed85dab25ce945c46726.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                16,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/1716740/ss_930710a45c08eaa4c10fa0be7c0663900e1d32f3.1920x1080.jpg?t=1687290127'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                2,
                'https://youtube.com/embed/xtzsuQgSPsU'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_820668033c382d05c51041e6e6470ab3bea94bba.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_d3f86d61086b91e1692d147958f513ac352c442e.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_71657db048cd5d90d639b1b6a75a3745f6aa019f.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_714b3857056522f28f3e617572956f63cc8623eb.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_69399a5322a79991c2e3e7abf41cbb78c042fc29.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_39590026d41de69bceaf1c27205e7383896be752.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_c3b0e5371b789ad4c2a061e42d3e7aca2eb07575.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_ce4ab5b666e85c132f836c54a8c111575481972f.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_7ec729a97789f94869b6b51f4c0d9af574be3bc3.1920x1080.jpg?t=1684861231'
            )
        `, `
            INSERT INTO media (
                game_id, type, uri
            ) VALUES (
                17,
                1,
                'https://cdn.cloudflare.steamstatic.com/steam/apps/2139230/ss_086363fd9b91d3fc1b9a5ee762fa6bac4b991da2.1920x1080.jpg?t=1684861231'
            )
        `,
	}

	for _, q := range queries {
		if _, err := database.Context.Exec(q); err != nil {
			return errors.New(fmt.Sprintln(err.Error(), q))
		}
	}

	return nil
}
