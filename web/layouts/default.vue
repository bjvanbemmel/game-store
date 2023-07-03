<template>
    <section class="w-full flex flex-col md:flex-row">
        <ClientOnly>
            <DefaultModal
                :active="mobileSearch"
                @toggle="(v) => mobileSearch = v"
            >
                <NavigationSearch
                    class="w-64 sm:w-80 md:w-96"
                    :autofocus="true"
                    :blackbox="false"
                />
            </DefaultModal>
        </ClientOnly>
        <NavigationBar />
        <main class="relative w-full p-8">
            <slot></slot>
            <p
                class="md:hidden pb-16 m-4 text-sm text-zinc-500"
            >
                Need help?
                <NuxtLink
                    class="text-sm hover:text-zinc-600 hover:cursor-pointer transition-colors duration-75"
                    to="/contact"
                >
                    <u>Contact us.</u>
                </NuxtLink>
            </p>
        </main>
    </section>
</template>

<script setup lang="ts">
const mobileSearch: Ref<Boolean> = useMobileSearch()
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
