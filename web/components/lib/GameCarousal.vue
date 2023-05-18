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
    <div
        class="flex flex-wrap gap-5"
        v-else
    >
        <GameThumbnail
            v-for="game, i in props.games"
            :key="i"
            :game="game"
        />
    </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import Game from '~/types/game'

const props = defineProps<{
    games: Array<Game> | null,
}>()
</script>
