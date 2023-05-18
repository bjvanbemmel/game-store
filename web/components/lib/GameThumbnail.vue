<template>
    <div
        class="group rounded-md w-48 hover:cursor-pointer"
    >
        <img
            :src="props.game.Thumbnail"
            class="h-60 w-full rounded-t-md object-cover"
        />
        <div class="flex flex-col transition-colors duration-100 gap-0.5 bg-zinc-800 group-hover:bg-zinc-700 px-3 py-2 overflow-clip shadow-md rounded-b-md line-clamp-1">
            <h1
                class="h-6 line-clamp-1"
            >
                {{ props.game.Title }}
            </h1>
            <a
                href=""
                class="h-5 w-max line-clamp-1 transition-colors duration-100 text-sm text-zinc-500 group-hover:text-zinc-400 hover:underline"
            >
                {{ props.game.Developer }}
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
