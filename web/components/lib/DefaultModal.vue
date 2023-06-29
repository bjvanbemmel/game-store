<template>
    <Transition>
        <div
            @click.stop="toggle"
            v-if="active"
            class="fixed top-0 left-0 z-40 bg-black/50 w-screen h-screen flex justify-center items-center cursor-pointer"
        >
            <DefaultContainer
                @click.stop
                class="flex flex-col gap-4 cursor-auto z-50"
            >
                <slot />
            </DefaultContainer>
        </div>
    </Transition>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, watch } from 'vue'

const emit = defineEmits<{
    (e: 'toggle', value: boolean): void,
}>()

function toggle() {
    emit('toggle', !props.active)
}

function toggleOnKey(e: KeyboardEvent) {
    if (e.key === 'Escape') {
        toggle()
    }
}

const props = defineProps({
    active: {
        type: Boolean,
        required: false,
        default: false,
    },
})

watch(props, (props) => {
    if (props.active) {
        window.document.body.style.overflow = 'hidden'
        document.addEventListener('keydown', toggleOnKey)
        return
    }

    document.removeEventListener('keydown', toggleOnKey)
    window.document.body.style.overflow = 'auto'
}, { immediate: true })
</script>

<style scoped>
.v-enter-active,
.v-leave-active {
    transition: opacity 0.2s ease;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
}
</style>
