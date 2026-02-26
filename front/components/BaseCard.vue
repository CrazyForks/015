<script setup lang="ts">
const props = defineProps<{
    title?: string
}>()
</script>

<template>
    <div class="rounded-xl p-5 bg-white/50 backdrop-blur-xl w-full lg:w-200">
        <div v-if="title" class="text-xl font-normal">{{ title }}</div>
        <slot />
    </div>
</template>
