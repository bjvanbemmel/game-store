<template>
    <div class="w-20 md:w-96 flex-shrink-0 transition-all duration-100"></div>
    <nav class="flex flex-col justify-between w-20 md:w-96 flex-shrink-0 bg-zinc-800 h-screen fixed z-20 transition-all duration-100">
        <ul class="flex flex-col gap-3 text-zinc-200 p-3 md:p-8 w-full">
            <NavigationSearch />
            <NavigationLink
                route=""
                class="h-[3.75rem] flex md:hidden justify-start md:justify-center items-center"
            >
                <MagnifyingGlassIcon
                    class="h-6"
                />
                <span class="hidden md:block">Most popular</span>
            </NavigationLink>
            <NavigationLink
                route="/"
                class="h-[3.75rem] flex justify-start md:justify-center items-center"
            >
                <ChartBarIcon
                    class="h-6"
                />
                <span class="hidden md:block">Most popular</span>
            </NavigationLink>
            <NavigationLink
                route="/most-loved"
                class="h-[3.75rem] flex items-center"
            >
                <HeartIcon
                    class="h-6"
                />
                <span class="hidden md:block">Most loved</span>
            </NavigationLink>
            <NavigationLink
                route="/upcoming-titles"
                class="h-[3.75rem] flex items-center"
            >
                <CalendarIcon
                    class="h-6"
                />
                <span class="hidden md:block">Upcoming titles</span>
            </NavigationLink>
        </ul>

        <NuxtLink
            class="m-4 text-sm text-zinc-500 hover:text-zinc-600 hover:cursor-pointer transition-colors duration-75 w-max"
            to="/contact"
        >
            Need help? <u>Contact us.</u>
        </NuxtLink>
    </nav>
</template>

<script setup lang="ts">
import { ChartBarIcon, HeartIcon, CalendarIcon, MagnifyingGlassIcon} from '@heroicons/vue/24/outline'
</script>
