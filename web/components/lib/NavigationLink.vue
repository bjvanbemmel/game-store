<template>
    <li
        class="hover:bg-zinc-700  rounded-md w-full"
        :class="{
            'bg-zinc-900': props.route === route.path,
        }"
    >
        <NuxtLink
            class="flex md:justify-start justify-center items-center gap-5 p-4 w-full text-lg"
            :to="props.route"
        >
            <slot />
        </NuxtLink>
    </li>
</template>

<script setup lang="ts">
const route = useRoute()

const props = defineProps<{
    route: string,
}>()
</script>
