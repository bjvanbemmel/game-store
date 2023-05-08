<template>
    <div
        class="group rounded-md w-48 hover:cursor-pointer"
    >
        <div
            v-if="game.id === undefined"
            class="h-40 w-full rounded-t-md border-none bg-zinc-600 animate-pulse"
        ></div>
        <img
            v-else
            :src="game.thumbnail"
            class="h-40 w-full rounded-t-md object-cover"
        />
        <div class="flex flex-col transition-colors duration-100 gap-0.5 bg-zinc-800 group-hover:bg-zinc-700 px-3 py-2 overflow-clip shadow-md rounded-b-md line-clamp-1">
            <div
                v-if="game.id === undefined"
                class="h-4 my-1 bg-zinc-600 rounded-md animate-pulse"
            ></div>
            <h1
                v-else
                class="h-6 line-clamp-1 empty:bg-zinc-600 empty:rounded-md empty:animate-pulse"
            >
                {{ props.game.title }}
            </h1>
            <div
                v-if="game.id === undefined"
                class="h-3 my-1 bg-zinc-600 rounded-md animate-pulse"
            ></div>
            <a
                v-else
                href=""
                class="h-5 w-max line-clamp-1 transition-colors duration-100 text-sm text-zinc-500 group-hover:text-zinc-400 hover:underline"
            >
                {{ props.game.developer }}
            </a>
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, PropType } from 'vue'
import Game from 'types/game'

const props = defineProps({
    game: {
        type: Object as PropType<Game>,
        required: true,
    },
})
</script>
