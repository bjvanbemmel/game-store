<template>
    <iframe
        v-if="media.type === MediaType.YouTube"
        class="rounded-md"
        :src="media.uri"
        title="YouTube video player"
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
        allowfullscreen
    />
    <img
        v-else
        :src="media.uri"
        class="rounded-md object-scale-down bg-zinc-900"
    />
</template>

<script setup lang="ts">
import Media from '~/types/media'
import MediaType from '~/types/media_type'
import { defineProps } from 'vue'

defineProps<{
    media: Media
}>()
</script>
