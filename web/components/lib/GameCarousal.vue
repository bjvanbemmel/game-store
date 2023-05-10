<template>
    <div
        class="flex flex-wrap gap-5"
    >
        <SkeletonGameThumbnail
            v-for="i in 14"
            :key="i"
        />
    </div>
</template>
