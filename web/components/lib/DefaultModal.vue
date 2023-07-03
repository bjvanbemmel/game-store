<template>
    <Transition>
        <div
            @click.stop="toggle"
            v-if="active"
            class="fixed top-0 left-0 z-40 p-8 8 bg-black/70 w-screen h-screen flex cursor-pointer"
            :class="centered ? 'justify-center items-center' : ''"
        >
            <DefaultContainer
                @click.stop
                class="flex flex-col gap-4 cursor-auto z-50 h-max"
            >
                <slot />
            </DefaultContainer>
        </div>
    </Transition>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, watch } from 'vue'

const props = defineProps({
    active: {
        type: Boolean,
        required: true,
    },
    centered: {
        type: Boolean,
        default: true,
    },
})

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

watch(
    () => props.active,
    (active) => {
        if (active) {
            document.body.style.overflow = 'hidden'
            document.addEventListener('keydown', toggleOnKey)
            return
        }

        document.removeEventListener('keydown', toggleOnKey)
        document.body.style.overflow = 'auto'
    }, { immediate: true }
)
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
