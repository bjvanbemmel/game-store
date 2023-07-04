<template>
    <section class="w-full flex flex-col md:flex-row">
        <ClientOnly>
            <DefaultModal
                :active="mobileSearch"
                @toggle="(v) => mobileSearch = v"
                :centered="false"
                class="justify-center"
            >
                <NavigationSearch
                    class="w-72 sm:w-80 md:w-96"
                    :autofocus="true"
                    :blackbox="false"
                />
            </DefaultModal>
        </ClientOnly>
        <NavigationBar />
        <main class="relative w-full px-2 py-4 md:px-8 md:py-8">
            <slot></slot>
            <p
                class="md:hidden pb-16 m-4 text-sm text-zinc-500"
                v-if="router.currentRoute.value.path !== '/contact'"
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
const mobileSearch: Ref<boolean> = useMobileSearch()
const router = useRouter()
</script>
