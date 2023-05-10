<template>
    <div
        ref="self"
        class="flex flex-wrap gap-4"
    >
        <SkeletonGameThumbnail
            v-for="i in data.thumbnailQuantity"
            :key="i"
        />
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

const self = ref<HTMLCanvasElement | null>(null)
const data = reactive({
    thumbnailQuantity: <number>7,
})
</script>
