<template>
    <section class="w-full flex flex-row">
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
