<template>
    <div class="flex flex-wrap gap-4">
        <section
            v-for="date, i in dateStrings"
            :key="i"
            class="flex flex-col gap-4 bg-black/30 p-4 rounded-md"
        >
            <h1
                class="flex gap-2 font-semibold items-center"
            >
                <CalendarIcon
                    class="h-6"
                />
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
import { CalendarIcon } from '@heroicons/vue/24/outline'

const games: Ref<Array<Game> | null> = ref(null)
const dateStrings: ComputedRef<Array<Date> | undefined> = computed(() => {
    return games.value?.map(x => x.release_date)
})

function parseDate(date: Date): String
{
    let d: Date = new Date(date)
    return d.toLocaleDateString('en-CA')
}

await useApi<Array<Game>>('/api/games', {
    onResponse({ response }) {
        if (response.status !== 200) return
        games.value = response._data
    },
    onResponseError({ response }) {
        console.error(response)
    },
    query: {
        filterBy: 'upcoming',
    },
    server: false,
})
</script>
