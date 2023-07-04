<template>
    <ClientOnly>
        <DefaultModal
            :active="data.modal"
            @toggle="(v) => data.modal = v"
        >
            <div class="relative group">
                <GameMedia
                    :media="data.activeMedia"
                    class="object-none w-96 md:w-[30rem] lg:w-[51rem] 2xl:w-[59rem]"
                />
                <div class="hidden group-hover:flex absolute h-full w-full top-0 items-center justify-between p-4">
                    <button
                        @click.stop="navigateMedia(Direction.Left)"
                        class="h-12 w-12"
                    >
                        <ChevronLeftIcon class="h-12 w-12 bg-zinc-700 shadow-md rounded-full p-3 hover:bg-zinc-600 transition-colors duration-75" />
                    </button>
                    <button
                        @click.stop="navigateMedia(Direction.Right)"
                        class="h-12 w-12"
                    >
                        <ChevronRightIcon class="h-12 w-12 bg-zinc-700 shadow-md rounded-full p-3 hover:bg-zinc-600 transition-colors duration-75" />
                    </button>
                </div>
            </div>
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
        class="w-full md:w-auto group relative cursor-pointer"
    >
        <GameMedia
            :media="data.activeMedia"
            class="h-64 md:h-[30rem] w-full md:w-[25rem] lg:w-[45rem] 2xl:w-[53rem]"
        />
        <MagnifyingGlassIcon
            v-if="data.activeMedia.type === MediaType.Image"
            class="h-12 absolute right-2 bottom-2 hidden group-hover:block bg-zinc-800 shadow-md rounded-full p-3"
        />
    </div>
    <div class="flex gap-2 w-full md:w-[25rem] lg:w-[45rem] 2xl:w-[53rem] overflow-x-scroll">
        <img
            v-for="md, i in media"
            :key="i"
            :src="getThumbnailIfYouTube(md)"
            @click.stop="data.activeMedia = md"
            class="rounded-md h-24 w-40 md:h-32 md:w-52 cursor-pointer object-cover"
            :class="md.uri === data.activeMedia.uri ? 'border-2 border-white' : ''"
        />
    </div>
</template>

<script setup lang="ts">
import { MagnifyingGlassIcon, ChevronLeftIcon, ChevronRightIcon } from '@heroicons/vue/20/solid'
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

let scrollThroughCarousal: NodeJS.Timer

watch(
    () => data.modal,
    (newVal: boolean) => {
        if (newVal) {
            clearInterval(scrollThroughCarousal)
            window.addEventListener('keydown', navigateUsingKeys)
            return
        }

        scrollThroughCarousal = setInterval(() => {
            navigateMedia(Direction.Right)
        }, 5000)
        window.removeEventListener('keydown', navigateUsingKeys)
    },
)

watch(
    () => data.activeMedia,
    (newVal: Media, oldVal: Media | undefined) => {
        if (newVal.type === oldVal?.type) return

        if (newVal.type === MediaType.YouTube) {
            clearInterval(scrollThroughCarousal)
            return
        }

        scrollThroughCarousal = setInterval(() => {
            navigateMedia(Direction.Right)
        }, 5000)
    }, { immediate: true }
)
</script>
