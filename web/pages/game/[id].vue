<template>
    <div class="flex flex-col gap-4">
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

        <DefaultContainer>
            <h1
                class="font-semibold"
            >
                Player count last 24 hours:
            </h1>
            <GamePlayerCountChart
                v-if="game"
                :game="game"
            />
        </DefaultContainer>
    </div>
</template>

<script setup lang="ts">
import Game from '~/types/game'

const route = useRoute()
const game: Ref<Game | null> = ref(null)

await useApi<Game>(`/api/games/${route.params.id}`, {
    onResponse({ response }) {
        game.value = response._data
    },
    server: false,
})
</script>
