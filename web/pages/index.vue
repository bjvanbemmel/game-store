<template>
    <div class="flex flex-col items-center md:items-start flex-wrap gap-4 w-full">
        <h1
            class="text-xl"
        >
            All games sorted by <b>popularity</b>
        </h1>

        <GameCarousal
            :games="games"
        ></GameCarousal>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Game from '~/types/game'

const games: Ref<Array<Game> | null> = ref(null)

await useApi<Array<Game>>('/api/games', {
    onResponse({ response }) {
        if (response.status !== 200) return
        games.value = response._data
    },
    onResponseError({ response }) {
        console.error(response)
    },
    query: {
        sortBy: 'popularity',
    },
    server: false,
})
</script>
