<template>
    <div class="flex flex-col gap-4">
        <div
            class="bg-cyan-700 p-4 rounded-md shadow-md"
            v-if="game?.release_date && gameDate > new Date()"
        >
            This game has not been released yet. It's slated for <b>{{ formattedDate }}</b>.
        </div>

        <section class="flex gap-4">
            <DefaultContainer class="flex flex-col items-center gap-4 w-full">
                <GameMediaCarousal
                    v-if="game?.media"
                    :media="game.media"
                />
                <SkeletonGameMediaCarousal
                    v-else
                />
            </DefaultContainer>
            <GameDetailCard
                :game="game"
            />
        </section>

        <DefaultContainer
            class="flex flex-col gap-2"
        >
            <h1
                v-if="game"
                class="font-semibold"
            >
                Player count last 24 hours:
            </h1>
            <SkeletonContainer
                v-else
                class="h-6 w-52"
            />
            <GamePlayerCountChart
                v-if="game"
                :game="game"
            />
            <SkeletonContainer
                v-else
                class="h-64"
            />
        </DefaultContainer>
    </div>
</template>

<script setup lang="ts">
import Game from '~/types/game'
import { computed } from 'vue'

const route = useRoute()
const game: Ref<Game | null> = ref(null)

const gameDate: ComputedRef<Date> = computed(() => {
    return new Date(game.value?.release_date ?? '');
})

const formattedDate: ComputedRef<String> = computed(() => {
    let date: Date = new Date(game.value?.release_date ?? '')
    return date.toLocaleDateString('nl-NL')
})

await useApi<Game>(`/api/games/${route.params.id}`, {
    onResponse({ response }) {
        game.value = response._data
    },
    server: false,
})
</script>
