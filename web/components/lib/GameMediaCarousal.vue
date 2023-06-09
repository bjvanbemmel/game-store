<template>
    <ClientOnly>
        <DefaultModal
            :active="data.modal"
            @toggle="(v) => data.modal = v"
        >
            <GameMedia
                :media="data.activeMedia"
                class="object-none w-auto max-w-7xl"
            />
            <p class="w-full text-center">{{ activeMediaId }} / {{ data.iterableMedia.length }}</p>
            <DefaultButton
                @click.stop="data.modal = !data.modal"
            >
                Close
            </DefaultButton>
        </DefaultModal>
    </ClientOnly>
    <div
        @click.stop="data.modal = !data.modal"
        class="group relative cursor-pointer"
    >
        <GameMedia
            :media="data.activeMedia"
            class="h-[30rem] w-[53rem]"
        />
        <MagnifyingGlassIcon
            v-if="data.activeMedia.type === MediaType.Image"
            class="h-12 absolute right-2 bottom-2 hidden group-hover:block bg-zinc-800 shadow-md rounded-full p-3"
        />
    </div>
    <div class="flex gap-2 w-[53rem] overflow-x-scroll">
        <img
            v-for="md, i in media"
            :key="i"
            :src="getThumbnailIfYouTube(md)"
            @click.stop="data.activeMedia = md"
            class="rounded-md h-32 w-52 cursor-pointer object-cover"
        />
    </div>
</template>

<script setup lang="ts">
import { MagnifyingGlassIcon } from '@heroicons/vue/20/solid'
import { reactive, computed, defineProps, watch } from 'vue'
import Media from '~/types/media'
import MediaType from '~/types/media_type'

const props = defineProps<{
    media: Array<Media>
}>()

const data = reactive<{
    activeMedia: Media,
    iterableMedia: Array<Media>,
    modal: boolean,
}>({
    activeMedia: props.media[0],
    iterableMedia: props.media.filter(x => x.type === MediaType.Image),
    modal: false,
})

function getThumbnailIfYouTube(media: Media): string
{
    if (media.type !== MediaType.YouTube) {
        return media.uri
    }

    let code: string = /^https:\/\/youtube\.com\/embed\/([A-z0-9]+)/.exec(media.uri)?.[1] ?? ''

    if (code === '') throw new Error('Could not get code from URI.')

    return `https://img.youtube.com/vi/${code}/0.jpg`
}

enum Direction {
    Left,
    Right,
}

function navigateUsingKeys(e: KeyboardEvent)
{
    if (data.iterableMedia.length === 1) return

    switch (e.key) {
        case "ArrowLeft":
            navigateMedia(Direction.Left)
            break
        case "ArrowRight":
            navigateMedia(Direction.Right)
            break
    }
}

function navigateMedia(direction: Direction)
{
    if (data.iterableMedia.length === 1) return

    switch (direction) {
        case Direction.Left:
            if (data.activeMedia.uri === data.iterableMedia[0].uri) {
                data.activeMedia = props.media[props.media.length - 1]
                break
            }

            data.activeMedia = props.media[activeMediaId.value - 1]
            break
        case Direction.Right:
            if (data.iterableMedia[data.iterableMedia.length - 1].uri === data.activeMedia.uri) {
                data.activeMedia = data.iterableMedia[0]
                break
            }
            
            data.activeMedia = props.media[activeMediaId.value + 1]
            break
    }
}

const activeMediaId: ComputedRef<number> = computed(() => {
    for (const [i, media] of props.media.entries()) {
        if (media.uri !== data.activeMedia.uri) continue

        return i
    }

    return 0
})

watch(data, () => {
    if (data.modal) {
        window.addEventListener('keydown', navigateUsingKeys)
        return
    }

    window.removeEventListener('keydown', navigateUsingKeys)
})
</script>
