<script setup lang="ts">
import AboutBaseInfo from '@/components/About/AboutBaseInfo.vue'
import AboutVersionView from '@/components/About/AboutVersionView.vue'

const { t } = useI18n()
</script>

<template>
    <BaseCard class="my-5 flex flex-col gap-5" :title="t('page.about.title')">
        <AboutBaseInfo />
        <AboutChartView />
        <AboutVersionView />
    </BaseCard>
</template>
