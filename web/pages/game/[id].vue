<template>
    <div class="flex flex-col gap-4">
        <section class="flex gap-4">
            <DefaultContainer class="flex flex-col items-center gap-4 w-full">
                <GameMediaCarousal
                    :media="media"
                />
            </DefaultContainer>
            <GameDetailCard
                :game="game"
            />
        </section>

        <DefaultContainer>
        </DefaultContainer>
    </div>
</template>

<script setup lang="ts">
import Game from '~/types/game'
import Image from '~/types/image'
import YouTube from '~/types/youtube'
import Media from '~/types/media'

const route = useRoute()
const game: Ref<Game | null> = ref(null)

// temp mock data
const media: Array<Media> = [
    new YouTube('https://youtube.com/embed/h542YbZuwkQ'),
    new Image('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftwinfinite.net%2Fwp-content%2Fuploads%2F2018%2F12%2FFFXIV-4-3.jpg&f=1&nofb=1&ipt=6942573503379020c50aa9a87c70128d2b92ad76cdfb39f9663cecf9999ef01f&ipo=images'),
    new Image('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.unpause.asia%2Fwp-content%2Fuploads%2Ffinal-fantasy-xiv-races-classes-jobs-square-enix.jpg&f=1&nofb=1&ipt=eb514e1c10f23ce289b1103b9c795b21f122d211293a249e40092d95ec408e3c&ipo=images'),
    new Image('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsteamcdn-a.akamaihd.net%2Fsteam%2Fapps%2F39210%2Fss_ddb71dc0c9a667de5c23e0b3823dad3590d5e9bb.1920x1080.jpg%3Ft%3D1549473655&f=1&nofb=1&ipt=c5fd0313fb76460180aaa65747ec6a6b21bc69aa5ac9a91e2fe2da20f6c94b95&ipo=images'),
    new Image('https://external-content.duckduckgo.com/iu/?u=http%3A%2F%2Fthehiddenlevels.com%2Fwp-content%2Fuploads%2F2017%2F05%2FFFXIV_40_Mediakit_20170501_020.jpg&f=1&nofb=1&ipt=8d75c93db281dcaf1c80717cae5c266af57195aed588273cd787cc82538f86f7&ipo=images'),
    new Image('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi1.wp.com%2Ftwinfinite.net%2Fwp-content%2Fuploads%2F2020%2F11%2Fffxiv-nier.jpg%3Ffit%3D1920%252C1080%26ssl%3D1&f=1&nofb=1&ipt=554f2a7ea743d17362ff922bbda1a695ad4f934c90ace85607efdc4723d43dcc&ipo=images'),
    new Image('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.vlznhe_8dyPLfT2WwVekBgHaJm%26pid%3DApi&f=1&ipt=a66a8f26c7ff79279dc956719eba3f3ea2dc119942c444f16680c1123a2fdea9&ipo=images'),
]

await useApi<Game>(`/api/games/${route.params.id}`, {
    onResponse({ response }) {
        game.value = response._data
    },
    server: false,
})
</script>
