<template>
    <li
        class="hover:bg-zinc-700 rounded-md w-full"
        :class="{
            'bg-zinc-900': props.route === route.name,
        }"
    >
        <NuxtLink
            class="flex items-center gap-5 p-4 w-full text-lg"
            to="/"
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
