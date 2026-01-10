<script setup lang="ts">
import getFileSize from '~/lib/getFileSize'
import { clamp } from 'lodash-es'
import FileUploadBlockProgressView from '@/components/FileUploadBlockProgressView.vue'
import FileUploadHeatMapView from '@/components/FileUploadHeatMapView.vue'
import type { SelectedFile, Uploadfile } from './types'

const { t } = useI18n()

const props = defineProps<{
    uploadfiles: Uploadfile[]
    selectedFile: SelectedFile
}>()
const selectedUploadfile = computed(() => props.uploadfiles.find((item) => item.fileId === props.selectedFile))
const selectedUploadfileChunk = computed(() => Object.values(selectedUploadfile.value?.uploadInfo?.chunks || {}))
const selectedUploadfileViewMode = ref<'progress' | 'heatmap'>('progress')
</script>

<template>
    <div class="col-span-4 flex flex-col bg-white/80 rounded-xl p-3 gap-5" v-if="selectedFile">
        <div class="flex flex-row gap-2 items-center">
            <LucideInfo class="size-4" />
            {{ t('page.progress.file.uploadDetails') }}
        </div>
        <div class="grid grid-cols-3 text-sm gap-3">
            <div>
                {{ t('page.progress.file.chunk') }}: {{ selectedUploadfile?.uploadInfo?.chunkLength }} x
                {{ getFileSize(selectedUploadfile?.uploadInfo?.ChunkSize as number) }}
            </div>
            <div class="truncate col-span-2">Hash: {{ selectedUploadfile?.hash }}</div>
            <div>{{ t('page.progress.file.completed') }}: {{ selectedUploadfileChunk?.filter((r) => r.status === 'success')?.length || 0 }}</div>
            <div>{{ t('page.progress.file.discarded') }}: {{ selectedUploadfileChunk?.filter((r) => r.status === 'error')?.length || 0 }}</div>
            <div>
                {{ t('page.progress.file.pending') }}:
                {{ (selectedUploadfile?.uploadInfo?.chunkLength || 0) - (selectedUploadfileChunk?.length || 0) }}
            </div>
            <div class="col-span-3 flex flex-row justify-between items-center">
                <div class="text-md font-bold">
                    {{ selectedUploadfileViewMode === 'progress' ? t('page.progress.file.chunkProgress') : t('page.progress.file.chunkHeatmap') }}
                </div>
                <Button
                    size="sm"
                    class="ml-auto text-xs"
                    @click="selectedUploadfileViewMode = selectedUploadfileViewMode === 'progress' ? 'heatmap' : 'progress'"
                >
                    <LucideArrowDownUp class="size-4" />
                    {{ selectedUploadfileViewMode === 'progress' ? t('page.progress.file.heatmap') : t('page.progress.file.progressBar') }}
                </Button>
            </div>
            <div class="h-7 col-span-3 flex flex-row gap-2 items-center" v-if="selectedUploadfileViewMode === 'progress'">
                <div class="flex-1 h-full">
                    <FileUploadBlockProgressView :data="selectedUploadfile?.uploadInfo" />
                </div>
                {{ clamp(((selectedUploadfileChunk?.length || 0) / (selectedUploadfile?.uploadInfo?.chunkLength || 0)) * 100, 0, 100)?.toFixed(2) }}%
            </div>
            <div v-if="selectedUploadfileViewMode === 'heatmap'" class="col-span-3 bg-black/5 rounded p-2">
                <FileUploadHeatMapView :size="12" :data="selectedUploadfile?.uploadInfo" />
            </div>
        </div>
    </div>
</template>
