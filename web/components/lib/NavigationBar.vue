<template>
    <div class="w-20 lg:w-96 flex-shrink-0 transition-all duration-100"></div>
    <nav class="flex flex-col justify-between w-20 lg:w-96 flex-shrink-0 bg-zinc-800 h-screen fixed z-20 transition-all duration-100">
        <ul class="flex flex-col gap-3 text-zinc-200 p-3 lg:p-8 w-full">
            <NavigationSearch
                class="hidden lg:flex mb-8"
            />
            <NavigationLink
                route=""
                class="h-[3.75rem] flex lg:hidden justify-start lg:justify-center items-center hover:cursor-pointer"
                @click="toggleMobileSearch"
            >
                <MagnifyingGlassIcon
                    class="h-6"
                />
                <span class="hidden lg:block">Most popular</span>
            </NavigationLink>
            <NavigationLink
                route="/"
                class="h-[3.75rem] flex justify-start lg:justify-center items-center"
            >
                <ChartBarIcon
                    class="h-6"
                />
                <span class="hidden lg:block">Most popular</span>
            </NavigationLink>
            <NavigationLink
                route="/most-loved"
                class="h-[3.75rem] flex items-center"
            >
                <HeartIcon
                    class="h-6"
                />
                <span class="hidden lg:block">Most loved</span>
            </NavigationLink>
            <NavigationLink
                route="/upcoming-titles"
                class="h-[3.75rem] flex items-center"
            >
                <CalendarIcon
                    class="h-6"
                />
                <span class="hidden lg:block">Upcoming titles</span>
            </NavigationLink>
        </ul>

        <p
            class="m-4 text-sm text-zinc-500"
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
