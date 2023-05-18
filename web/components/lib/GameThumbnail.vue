<template>
    <NuxtLink
        @mouseenter="setActiveId(props.game.ID)"
        @mouseleave="setActiveId(-1)"
        to="/"
        :title="props.game.Title"
        :class="{
            'opacity-60': ![-1, props.game.ID, ].includes(props.activeId),
        }"
        class="group transition-opacity duration-150 rounded-md w-48 hover:cursor-pointer"
    >
        <img
            :src="props.game.Thumbnail"
            class="h-60 w-full rounded-t-md object-cover"
        />
        <div class="flex flex-col gap-0.5 bg-zinc-800 px-3 py-2 overflow-clip shadow-md rounded-b-md line-clamp-1">
            <h1
                class="h-6 font-semibold line-clamp-1"
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
    </NuxtLink>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, PropType } from 'vue'
import Game from 'types/game'

const props = defineProps({
    game: {
        type: Object as PropType<Game>,
        required: true,
    },
    activeId: {
        type: Number,
        required: true,
    },
})

const emit = defineEmits<{
    (e: 'active', value: number): void
}>()

function setActiveId(id: number) {
    emit('active', id)
}
</script>
