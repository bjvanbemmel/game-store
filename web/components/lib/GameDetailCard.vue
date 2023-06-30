<template>
    <DefaultContainer class="flex flex-col gap-4">
        <!-- Cover -->
        <img
            v-if="game?.thumbnail"
            class="h-96 w-64 rounded-md object-cover outline-none border-none"
            :src="game?.thumbnail"
        />
        <SkeletonContainer
            v-else
            class="h-96 w-64"
        />

        <!-- Title -->
        <h1
            v-if="game?.title"
            class="text-lg font-semibold"
        >
            {{ game?.title }}
        </h1>
        <SkeletonContainer
            v-else
            class="h-9"
        />

        <!-- Rating -->
        <GameRatingLabel
            v-if="game"
            :rating="game.rating"
        />
        <SkeletonContainer
            v-else
            class="h-6 w-12"
        />

        <!-- Description -->
        <p
            v-if="game?.description"
            class="w-64 text-sm text-zinc-200"
        >
            {{ game.description }}
        </p>
        <SkeletonContainer
            v-else
            class="h-12"
        />

        <!-- Release date -->
        <p
            class="text-sm text-zinc-400"
            v-if="game?.release_date"
        >
            Release date: <span class="text-zinc-300">{{ gameDate.toLocaleDateString('nl-NL') }}</span>
        </p>
        <SkeletonContainer
            v-else
            class="h-6"
        />

        <!-- Developers -->
        <div
            v-if="game?.developers"
            class="flex flex-col gap-1"
        >
            <NuxtLink
                v-for="dev, i in game.developers"
                :key="i"
                class="text-sm text-zinc-400 hover:underline"
                :to="`/developer/${dev.id}`"
            >
                {{ dev.name }}
            </NuxtLink>
        </div>
        <SkeletonContainer
            v-else
            class="h-6 w-64"
        />

        <!-- Genres -->
        <div
            v-if="game?.genres"
            class="flex flex-wrap gap-2 w-64"
        >
            <NuxtLink
                v-for="genre, i in game.genres"
                :key="i"
                class="text-xs py-1 px-2 bg-zinc-900 border-[0.025rem] border-zinc-500 rounded-md hover:bg-zinc-700"
                :to="`/genre/${genre.id}`"
            >
                {{ genre.name }}
            </NuxtLink>
        </div>
        <div
            v-else
            class="flex flex-wrap gap-2 w-64"
        >
            <SkeletonContainer class="h-6 w-14" />
            <SkeletonContainer class="h-6 w-14" />
            <SkeletonContainer class="h-6 w-14" />
            <SkeletonContainer class="h-6 w-14" />
        </div>

        <!-- Purchase -->
        <div
            v-if="game"
        >
            <div
                v-if="game && game.price > 0"
                class="flex w-full"
            >
                <DefaultButton
                    class="rounded-r-none w-3/5"
                >
                    Add to cart
                </DefaultButton>
                <div
                    class="flex p-1 items-center justify-center bg-zinc-900 border-zinc-700 border rounded-r-md w-2/5"
                >
                    &euro; {{ game?.price }}
                </div>
            </div>
            <div
                v-else
                class="flex w-full"
            >
                <DefaultButton
                    class="rounded-r-none w-3/5"
                    :disabled="gameDate > new Date()"
                >
                    Add to library
                </DefaultButton>
                <div
                    class="flex p-1 items-center justify-center bg-zinc-900 border-zinc-700 border rounded-r-md w-2/5"
                >
                    FREE
                </div>
            </div>
        </div>
        <SkeletonContainer
            v-else
            class="h-12"
        />
    </DefaultContainer>
</template>

<script setup lang="ts">
import Game from '~/types/game'
import { defineProps, computed } from 'vue'

const props = defineProps<{
    game: Game | null
}>()

const gameDate: ComputedRef<Date> = computed(() => {
    return new Date(props.game?.release_date ?? '');
})
</script>
