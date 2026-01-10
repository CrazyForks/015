<script setup lang="ts">
import type { Uploadfile } from './types'

const { t } = useI18n()

const { uploadfiles } = defineProps<{
    uploadfiles: Uploadfile[]
}>()
const totalTaskStatus = computed(() => {
    if (uploadfiles.some((r) => r.status === 'start')) {
        return 'start'
    }
    if (uploadfiles.some((r) => r.status === 'pause')) {
        return 'pause'
    }
    return 'disabled'
})
const totalUploadProgress = computed(() => {
    const successCount = uploadfiles.reduce((acc, curr) => {
        const { status, uploadInfo } = curr || {}
        if (status === 'finish') return acc
        const { chunks } = uploadInfo || {}
        return acc + Object.entries(chunks || {}).filter(([index, chunk]) => chunk.status === 'success').length
    }, 0)
    const totalCount = uploadfiles.reduce((acc, curr) => {
        const { status, uploadInfo } = curr || {}
        if (status === 'finish') return acc
        const { chunkLength } = uploadInfo || {}
        return acc + (chunkLength || 0)
    }, 0)
    return ((successCount || 0) / (totalCount || 0)) * 100
})
</script>

<template>
    <div class="rounded-xl col-span-4 md:col-span-1 bg-white/80 h-32 md:h-auto md:aspect-square flex flex-col gap-2 relative overflow-hidden">
        <div class="absolute top-0 left-0 w-full h-full z-[0] flex flex-col justify-end">
            <div class="w-full bg-green-100 border-t border-green-500" :style="`height: ${totalUploadProgress || 0}%`"></div>
        </div>
        <div class="flex flex-col gap-2 justify-between p-3 h-full relative z-[1]">
            <div class="flex flex-col gap-1">
                <div class="text-xs opacity-70">{{ t('page.progress.file.totalUploadProgress') }}</div>
                <div class="text-4xl font-bold">{{ (totalUploadProgress || 0).toFixed(1) }}%</div>
            </div>
            <div class="flex flex-row gap-2">
                <Button
                    class="aspect-square hover:text-white p-0 bg-green-200 hover:bg-green-300 text-green-500"
                    :disabled="['start', 'disabled'].includes(totalTaskStatus)"
                    @click="
                        () => {
                            uploadfiles.forEach((r) => {
                                r.status = 'start'
                            })
                        }
                    "
                >
                    <LucidePlay class="size-1/2" />
                </Button>
                <Button
                    class="aspect-square hover:text-white p-0 bg-orange-200 hover:bg-orange-300 text-orange-500"
                    :disabled="['pause', 'disabled'].includes(totalTaskStatus)"
                    @click="
                        () => {
                            uploadfiles.forEach((r) => {
                                r.status = 'pause'
                            })
                        }
                    "
                >
                    <LucideSquare class="size-1/2" />
                </Button>
                <!-- <Button class="aspect-square bg-blue-200 hover:bg-blue-300 text-blue-500 hover:text-white p-0">
                        <LucideSettings class="size-1/2" />
                    </Button> -->
            </div>
        </div>
    </div>
</template>
