// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    app: {
        head: {
            title: 'Game store',
        },
    },
    runtimeConfig: {
        public: {
            baseURL: 'http://localhost:81',
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
