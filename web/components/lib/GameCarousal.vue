<template>
    <div
        class="flex md:justify-start flex-wrap gap-5"
        v-if="props.games === null"
    >
        <SkeletonGameThumbnail
            v-for="i in 10"
            :key="i"
        />
    </div>
    <div
        class="flex justify-center md:justify-start flex-wrap gap-5"
        v-else
    >
        <GameThumbnail
            v-for="game, i in props.games"
            :key="i"
            :game="game"
            :active-id="data.activeId"
            @active="(e: number) => data.activeId = e"
        />
    </div>
</template>

<script setup lang="ts">
import { defineProps, reactive, } from 'vue'
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
</script>
