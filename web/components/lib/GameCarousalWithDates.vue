<template>
    <div
        class="flex flex-wrap gap-5"
        v-if="props.games === null"
    >
        <SkeletonGameThumbnail
            v-for="i in 10"
            :key="i"
        />
    </div>
    <section
        v-if="props.games"
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
        <div class="flex flex-wrap gap-4">
            <GameThumbnail
                v-for="game, i in props.games.filter(x => x.release_date === date)"
                :key="i"
                :game="game"
                :active-id="data.activeId"
                @active="(e: number) => data.activeId = e"
            />
        </div>
    </section>
</template>

<script setup lang="ts">
import { CalendarIcon } from '@heroicons/vue/24/outline'
import { defineProps, reactive, computed, } from 'vue'
import Game from '~/types/game'

interface Data {
    activeId: number,
}

const data: Data = reactive({
    activeId: -1,
})

const props = defineProps<{
    games: Array<Game> | null,
}>()

const dateStrings: ComputedRef<Array<Date> | undefined> = computed(() => {
    return [...new Set(props.games?.map(x => x.release_date))]
})

function parseDate(date: Date): String
{
    let d: Date = new Date(date)
    return d.toLocaleDateString('en-CA')
}
</script>
