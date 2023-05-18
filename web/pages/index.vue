<template>
    <div class="flex flex-col flex-wrap gap-4 w-full">
        <GameCarousal
            :games="games"
        ></GameCarousal>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Game from '~/types/game'

const games: Ref<Array<Game> | null> = ref(null)

await useFetch<Array<Game>>('http://192.168.2.15:81/api/games', {
    onResponse({ response }) {
        if (response.status !== 200) return
        games.value = response._data
    },
    onResponseError({ response }) {
        console.error(response)
    },
    server: false,
})
</script>
