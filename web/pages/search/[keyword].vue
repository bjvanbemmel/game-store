<template>
    <div class="flex flex-col flex-wrap gap-4 w-full">
        <h1
            class="text-xl"
        >
            Search results for <q class="font-semibold">{{ route.params.keyword }}</q>
        </h1>

        <GameCarousal
            :games="games"
        />

        <h2
            v-if="games?.length === 0"
        >
            No results found...
        </h2>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Game from '~/types/game'

const route = useRoute()
const games: Ref<Array<Game> | null> = ref(null)

await useApi<Array<Game>>(`/api/games?keyword=${route.params.keyword}`, {
    onResponse({ response }) {
        if (response.status !== 200) return
        games.value = response._data ?? new Array<Game>()
    },
    onResponseError({ response }) {
        console.error(response)
    },
    server: false,
})
</script>
