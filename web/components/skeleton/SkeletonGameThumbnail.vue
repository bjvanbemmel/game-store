<template>
    <div
        class="group rounded-md w-48"
    >
        <SkeletonContainer
            class="h-60 w-full rounded-b-none"
        />
        <div class="flex flex-col transition-colors duration-100 gap-0.5 bg-zinc-800 px-3 py-3 overflow-clip shadow-md rounded-b-md line-clamp-1">
            <SkeletonContainer
                class="h-4 my-1"
            />
            <SkeletonContainer
                class="h-3 my-1"
            />
            <div class="my-1 flex justify-between">
                <SkeletonContainer
                    class="h-5 w-12"
                />
                <SkeletonContainer
                    class="h-5 w-12"
                />
            </div>
        </div>
    </div>
</template>
