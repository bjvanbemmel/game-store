<template>
    <div class="flex flex-col gap-4">
        <section class="flex gap-4">
            <DefaultContainer class="flex flex-col items-center gap-4 w-full">
                <SkeletonImageCarousal />
            </DefaultContainer>
            <GameDetailCard
                :game="game"
            />
        </section>

        <DefaultContainer>
        </DefaultContainer>
    </div>
</template>

<script setup lang="ts">
import Game from 'types/game'

const route = useRoute()
const game: Ref<Game | null> = ref(null)

await useApi<Game>(`/api/games/${route.params.id}`, {
    onResponse({ response }) {
        game.value = response._data
    },
    server: false,
})
</script>
