<template>
    <p
        class="border-[.025rem] rounded-md w-max py-0.5 px-1 text-xs"
        :class="colorFromRating(rating)"
        v-if="rating >= 0"
    >
        {{rating}}%
    </p>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

defineProps<{
    rating: number,
}>()

function colorFromRating(rating: number): String
{
    if (rating > 90) return 'bg-green-800 border-green-500'

    if (rating > 70) return 'bg-lime-800 border-lime-500'

    if (rating > 50) return 'bg-yellow-800 border-yellow-500'

    if (rating > 40) return 'bg-orange-800 border-orange-500'

    return 'bg-red-700'
}
</script>
