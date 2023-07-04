<template>
    <div class="fixed md:static bottom-0 w-full h-20 md:w-20 2xl:w-96 flex-shrink-0 transition-all duration-100"></div>
    <nav class="bottom-0 flex flex-col justify-between w-full md:w-20 2xl:w-96 flex-shrink-0 bg-zinc-800 h-20 md:h-screen fixed z-30 transition-all duration-100">
        <ul class="flex justify-evenly md:justify-start md:flex-col gap-3 text-zinc-200 p-3 2xl:p-8 w-full">
            <NavigationSearch
                class="hidden 2xl:flex mb-8"
            />
            <NavigationLink
                route=""
                class="h-[3.75rem] flex 2xl:hidden justify-start 2xl:justify-center items-center hover:cursor-pointer"
                @click="toggleMobileSearch"
            >
                <MagnifyingGlassIcon
                    class="h-6"
                />
                <span class="hidden 2xl:block">Most popular</span>
            </NavigationLink>
            <NavigationLink
                route="/"
                class="h-[3.75rem] flex justify-start 2xl:justify-center items-center"
            >
                <ChartBarIcon
                    class="h-6"
                />
                <span class="hidden 2xl:block">Most popular</span>
            </NavigationLink>
            <NavigationLink
                route="/most-loved"
                class="h-[3.75rem] flex items-center"
            >
                <HeartIcon
                    class="h-6"
                />
                <span class="hidden 2xl:block">Most loved</span>
            </NavigationLink>
            <NavigationLink
                route="/upcoming-titles"
                class="h-[3.75rem] flex items-center"
            >
                <CalendarIcon
                    class="h-6"
                />
                <span class="hidden 2xl:block">Upcoming titles</span>
            </NavigationLink>
        </ul>

        <p
            class="hidden md:block m-4 text-sm text-zinc-500"
        >
            Need help?
            <NuxtLink
                class="text-sm hover:text-zinc-600 hover:cursor-pointer transition-colors duration-75"
                to="/contact"
            >
                <u>Contact us.</u>
            </NuxtLink>
        </p>
    </nav>
</template>

<script setup lang="ts">
import { ChartBarIcon, HeartIcon, CalendarIcon, MagnifyingGlassIcon} from '@heroicons/vue/24/outline'
</script>
