<template>
    <form
        class="group flex relative mb-8"
    >
        <input
            class="
            rounded-l-lg
            px-4
            py-2
            shadow-md
            w-5/6
            bg-zinc-600
            text-zinc-50
            focus:outline-none
            "
            type="text"
            placeholder="Search..."
        />
        <button
            class="bg-zinc-700 text-zinc-300 rounded-r-lg w-1/6"
        >
            <MagnifyingGlassIcon
                class="h-5 mx-auto"
            />
        </button>
    </form>
</template>

<script setup lang="ts">
import { MagnifyingGlassIcon } from '@heroicons/vue/20/solid';
</script>
