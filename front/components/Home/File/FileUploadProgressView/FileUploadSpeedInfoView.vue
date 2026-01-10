<script setup lang="ts">
const { t } = useI18n()
</script>
<template>
    <div class="flex flex-col gap-2 pt-3">
        <div class="text-xl font-bold">{{ t('page.progress.file.uploadSpeedInfo.title') }}</div>
        <div class="opacity-75">
            <i18n-t keypath="page.progress.file.uploadSpeedInfo.desc.base">
                <template #chunkNum>
                    <b>{{ t('page.progress.file.uploadSpeedInfo.desc.chunkNum') }}</b>
                </template>
                <template #chunkSize>
                    <b>{{ t('page.progress.file.uploadSpeedInfo.desc.chunkSize') }}</b>
                </template>
            </i18n-t>
        </div>
    </div>
</template>
