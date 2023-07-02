<template>
    <section class="w-full flex flex-row">
        <NavigationBar />
        <Transition>
            <div
                class="fixed lg:hidden top-0 z-30 p-5 flex w-screen justify-center items-center bg-zinc-900"
                v-if="mobileSearch"
            >
                <NavigationSearch
                    class="h-10 w-full"
                />
            </div>
        </Transition>
        <main class="relative w-full p-8">
            <slot></slot>
        </main>
    </section>
</template>

<script setup lang="ts">
const mobileSearch = useMobileSearch()
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
</style>
