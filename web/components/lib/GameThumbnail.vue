<template>
    <div
        class="relative h-[21.125rem] w-48"
    >
        <NuxtLink
            @mouseenter="setActiveId(props.game.id)"
            @mouseleave="setActiveId(-1)"
            :to="`/game/${props.game.id}`"
            :title="props.game.title"
            :class="{
                'opacity-60': ![-1, props.game.id, ].includes(props.activeId),
                'z-10': props.activeId === props.game.id,
            }"
            class="absolute group transition-opacity duration-150 rounded-md w-48 hover:cursor-pointer"
        >
            <div class="absolute h-60 w-48"></div>
            <img
                :src="props.game.thumbnail"
                class="h-60 w-full rounded-t-md object-cover"
            />
            <div class="flex flex-col gap-0.5 bg-zinc-800 px-3 py-3 overflow-clip shadow-md rounded-b-md line-clamp-1">
                <h1
                    class="h-6 font-semibold truncate"
                >
                    {{ props.game.title }}
                </h1>
                <NuxtLink
                    :to="`/developer/${developer.id}`"
                    :title="developer.name"
                    @click.stop
                    class="h-5 z-20 truncate transition-colors duration-100 text-sm text-zinc-500 group-hover:text-zinc-400 hover:underline"
                >
                    {{ developer.name }}
                </NuxtLink>
                <p
                    class="mt-1 border-zinc-500 border-[.025rem] rounded-md bg-zinc-900 w-max py-0.5 px-1 text-xs"
                    v-if="game.price > 0"
                >
                    &euro; {{game.price}}
                </p>
                <p
                    class="mt-1 border-zinc-500 border-[.025rem] rounded-md bg-zinc-900 w-max py-0.5 px-1 text-xs"
                    v-else
                >
                    FREE
                </p>
                <div
                    :class="props.activeId === props.game.id ? 'flex' : 'hidden'"
                    class="gap-1 mt-2 flex-wrap"
                >
                    <div
                        v-for="genre, i in props.game.genres"
                        :key="i"
                        class="border-zinc-500 border-[.025rem] rounded-md bg-zinc-900 py-0.5 px-1 text-xs"
                    >
                        {{ genre.name }}
                    </div>
                </div>
            </div>
        </NuxtLink>
    </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, PropType, computed } from 'vue'
import Game from '~/types/game'

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

const developer = computed(() => props.game.developers[0])

const emit = defineEmits<{
    (e: 'active', value: number): void
}>()

function setActiveId(id: number) {
    emit('active', id)
}
</script>
