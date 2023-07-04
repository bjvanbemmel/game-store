<template>
    <div class="flex flex-col items-center md:items-start flex-wrap gap-4 w-full">
        <h1
            v-if="genre"
            class="text-xl"
        >
            All <b>{{ genre.name }}</b> games
        </h1>

        <GameCarousal
            :games="games"
        ></GameCarousal>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Game from '~/types/game'
import Developer from '~/types/developer'

const route = useRoute()
const games: Ref<Array<Game> | null> = ref(null)
const genre: Ref<Developer | null> = ref(null)

await useApi<Array<Game>>(`/api/genres/${route.params.id}/games`, {
    onResponse({ response }) {
        if (response.status !== 200) return
        genre.value = response._data.genre
        games.value = response._data.games
    },
    onResponseError({ response }) {
        console.error(response)
    },
    server: false,
})
</script>
