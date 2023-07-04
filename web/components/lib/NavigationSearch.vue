<template>
    <div>
        <Transition>
            <div
                v-if="props.blackbox && active"
                class="absolute top-0 left-0 w-screen h-screen bg-black/40 z-40"
                @click.stop="active = false"
            ></div>
        </Transition>
        <form
            id="1"
            class="group w-full flex gap-4 flex-col relative"
            :class="active ? 'z-50' : ''"
            @keydown.enter.stop="searchIfNotEmpty(keyword)"
            @keydown.escape.stop="active = false"
            @submit.stop.prevent
        >
            <div
                class="relative flex item-center w-full" 
            >
                <div
                    class="relative flex w-5/6"
                >
                    <input
                        class="
                        rounded-l-md
                        px-4
                        py-2
                        shadow-md
                        w-full
                        bg-zinc-600
                        text-zinc-50
                        focus:outline-none
                        "
                        type="text"
                        placeholder="Search..."
                        v-model="keyword"
                        ref="search"
                        @click.stop="active = true"
                        @focus="active = true"
                        @input="active = true"
                    />

                    <Transition>
                        <button
                            v-if="active"
                            class="absolute right-2 h-full flex items-center"
                            @click="emptyAndFocus"
                        >
                            <XMarkIcon
                                class="h-6 text-zinc-400"
                            />
                        </button>
                    </Transition>
                </div>

                <button
                    class="bg-zinc-700 text-zinc-300 rounded-r-md w-1/6"
                    @click="searchIfNotEmpty(keyword)"
                >
                    <MagnifyingGlassIcon
                        class="h-5 mx-auto"
                    />
                </button>
            </div>

            <Transition
                name="games"
            >
                <NavigationSearchResults
                    v-if="active && keyword.length >= 2"
                    @navigate="active = false"
                    @browse="searchIfNotEmpty(keyword)"
                    :games="games"
                />
            </Transition>
        </form>
    </div>
</template>

<script setup lang="ts">
import { MagnifyingGlassIcon, XMarkIcon } from '@heroicons/vue/20/solid';
import Game from '~/types/game'
import { ref, watch, defineProps, onMounted } from 'vue'
import debounce from 'lodash/debounce';

const active: Ref<boolean> = ref(false);
const keyword: Ref<string> = ref("");
const games: Ref<Array<Game>> = ref(new Array<Game>());
const search: Ref<HTMLInputElement | null> = ref(null)
const router = useRouter()

const props = defineProps({
    blackbox: {
        type: Boolean,
        required: false,
        default: true,
    },
    autofocus: {
        type: Boolean,
        required: false,
        default: false,
    },
})

onMounted(() => {
    if (!props.autofocus) return

    search.value?.focus()
})

watch(keyword, debounce(async () => {
    if (keyword.value.length < 2) return

    await useApi<Array<Game>>(`/api/games?keyword=${keyword.value}`, {
        onResponse({ response }) {
            games.value = response._data ?? new Array<Game>()
        },
        onResponseError({ response }) {
            console.error(response)
        },
        server: false,
    })
}, 100))

function searchIfNotEmpty(keyword: string) {
    keyword = keyword.replaceAll('/', '')
    if (keyword === '') return

    active.value = false
    toggleMobileSearch()
    router.push(`/search/${keyword}`)
}

function emptyAndFocus() {
    keyword.value = ''
    search.value?.focus()
}

</script>

<style scoped>
.v-enter-from,
.v-leave-to {
    opacity: 0;
}

.v-enter-active,
.v-leave-active {
    transition: opacity 0.2s ease;
}

.games-enter-from,
.games-leave-to {
    opacity: 0;
    padding: 0px;
}

.games-enter-active,
.games-leave-active {
    transition: opacity 0.2s ease, padding 0.2s ease;
}
</style>
