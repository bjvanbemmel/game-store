<template>
    <div
        class="
        flex
        flex-col
        w-full
        bg-zinc-800
        rounded-md
        shadow-md
        absolute
        top-12
        p-2
        "
    >
        <div
            v-if="games.length === 0"
            class="p-2"
        >
            No results found...
        </div>
        <div
            v-else
            class="flex flex-col"
        >
            <NuxtLink
                v-for="game, i in games.slice(0, 4)"
                :key="i"
                @click.stop="emit('navigate')"
                :to="`/game/${game.id}`"
                class="
                first:rounded-t-md
                last:rounded-b-lg
                flex
                gap-3
                p-4
                items-center
                transition-colors
                duration-75
                hover:bg-zinc-700
                hover:cursor-pointer
                "
            >
                <img
                    :src="game.thumbnail"
                    class="
                    object-cover
                    w-14
                    h-10
                    "
                />
                <div
                    clas="flex flex-col"
                >
                    <h1
                        class="line-clamp-1"
                    >
                        {{ game.title }}
                    </h1>
                    <p
                        class="text-xs text-zinc-400"
                    >
                        &euro; {{ game.price }}
                    </p>
                </div>
            </NuxtLink>
            <NuxtLink
                class="rounded-b-md underline hover:bg-zinc-700 hover:cursor-pointer transition-colors duration-75 text-blue-400 p-4"
            >
                Browse all {{ games.length }} results
            </NuxtLink>
        </div>
    </div>
</template>

<script setup lang="ts">
import Game from '~/types/game'
import { defineProps, defineEmits } from 'vue'

defineProps<{
    games: Array<Game>
}>()

const emit = defineEmits<{
    (e: 'navigate'): void
}>()
</script>
