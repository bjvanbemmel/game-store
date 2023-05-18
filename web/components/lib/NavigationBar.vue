<template>
    <div class="w-96 flex-shrink-0"></div>
    <nav class="w-96 flex-shrink-0 bg-zinc-800 h-screen fixed z-20">
        <ul class="flex flex-col gap-3 text-zinc-200 p-8 w-full">
            <NavigationSearch />
            <NavigationLink
                route="/"
            >
                <ChartBarIcon
                    class="h-6"
                />
                Most popular
            </NavigationLink>
            <NavigationLink
                route="highest-rated"
            >
                <HeartIcon
                    class="h-6"
                />
                Most loved
            </NavigationLink>
            <NavigationLink
                route="upcoming"
            >
                <CalendarIcon
                    class="h-6"
                />
                Upcoming titles
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
import { ChartBarIcon, HeartIcon, BeakerIcon, CalendarIcon, } from '@heroicons/vue/24/outline'
</script>
