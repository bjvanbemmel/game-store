<template>
    <div class="flex flex-col flex-wrap gap-4 w-full">
        <h1
            class="text-xl"
        >
            Upcoming titles
        </h1>

        <section
            v-for="date, i in dateStrings"
            :key="i"
            class="flex flex-col gap-2"
        >
            <h1
                class="font-semibold text-xl"
            >
                {{ parseDate(date) }}:
            </h1>
            <GameCarousal
                v-if="games"
                :games="games.filter(x => x.release_date == date)"
            ></GameCarousal>
        </section>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Game from '~/types/game'

const games: Ref<Array<Game> | null> = ref(null)
const dateStrings: ComputedRef<Array<Date> | undefined> = computed(() => {
    return games.value?.map(x => x.release_date)
})

function parseDate(date: Date): String
{
    let d: Date = new Date(date)
    return d.toLocaleDateString('nl-NL')
}

await useApi<Array<Game>>('/api/games?filterBy=upcoming', {
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
