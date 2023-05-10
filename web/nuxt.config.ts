// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    app: {
        head: {
            title: 'Game store',
        },
    },
    components: [
        '~/components',
        '~/components/lib',
        '~/components/skeleton',
    ],
    css: [
        '~/assets/css/main.css',
    ],
    postcss: {
        plugins: {
            tailwindcss: {},
            autoprefixer: {},
        },
    },
})
