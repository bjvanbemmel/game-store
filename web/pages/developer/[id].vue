<template>
    <div class="flex flex-col flex-wrap gap-4 w-full">
        <h1
            class="text-xl"
        >
            All games from <b>{{ developer?.name }}</b>
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
const games: Ref<Array<Game>> = ref(new Array<Game>())
const developer: Ref<Developer | null> = ref(null)

await useApi<Array<Game>>(`/api/developers/${route.params.id}/games`, {
    onResponse({ response }) {
        if (response.status !== 200) return
        developer.value = response._data.developer
        games.value = response._data.games
    },
    onResponseError({ response }) {
        console.error(response)
    },
    server: false,
})
</script>
