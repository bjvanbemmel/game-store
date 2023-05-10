<template>
    <div class="w-96 flex-shrink-0"></div>
    <nav class="w-96 flex-shrink-0 bg-zinc-800 h-screen fixed z-20">
        <ul class="flex flex-col gap-3 text-zinc-200 p-8 w-full">
            <NavigationSearch />
            <NavigationLink
                route="index"
            >
                <ChartBarIcon
                    class="h-6"
                />
                Most popular
            </NavigationLink>
            <NavigationLink
                route="rating"
            >
                <HeartIcon
                    class="h-6"
                />
                Most loved
            </NavigationLink>
            <NavigationLink
                route="random"
            >
                <BeakerIcon
                    class="h-6"
                />
                Try your luck
            </NavigationLink>
        </ul>
    </nav>
</template>

<script setup lang="ts">
import { ChartBarIcon, HeartIcon, BeakerIcon, } from '@heroicons/vue/24/outline'
</script>
