<template>
    <div class="flex flex-wrap gap-4">
        <GameThumbnail
            v-for="game, i in data.games"
            :game="game"
            :key="i"
        />
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import Game from '../types/game'

const data = reactive({
    games: [
        <Game>{
            id: 0,
            title: "Final Fantasy XIV: A Realm Reborn",
            thumbnail: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fstatic2.thegamerimages.com%2Fwordpress%2Fwp-content%2Fuploads%2F2020%2F10%2FFinal-Fantasy-14-Chocobos-1.jpg&f=1&nofb=1&ipt=0a175f87625be88b90162c1aa535e00581f98395dadb9729dec648078a359bb3&ipo=images",
            developer: "Square Enix",
        },
        <Game>{},
        <Game>{},
        <Game>{},
        <Game>{},
    ],
})
</script>
