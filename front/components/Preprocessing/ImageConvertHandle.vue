<script setup lang="ts">
import SelectField from '../Field/SelectField.vue'
import FormButton from '../Field/FormButton.vue'
import type { FileShareHandleProps } from './types'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const props = defineProps<{
    hide: () => void
    file: File[]
    onFileHandle: (props: FileShareHandleProps) => void
}>()
</script>

<template>
    <VeeForm v-slot="{ values }" :initialValues="{ target_ext: 'jpg' }">
        <div class="flex flex-col gap-3">
            <h2 class="text-lg font-bold">{{ t('page.shareOptions.imageConvert.title') }}</h2>
            <div class="flex flex-row gap-2 items-center">
                <Label>{{ t('page.shareOptions.imageConvert.targetFormat') }}</Label>
                <SelectField
                    name="target_ext"
                    :label="t('page.shareOptions.imageConvert.targetFormat')"
                    :options="[
                        { label: 'JPG', value: 'jpg' },
                        { label: 'PNG', value: 'png' },
                        { label: 'WEBP', value: 'webp' },
                    ]"
                    rules="required"
                />
            </div>
            <FormButton
                @click="
                    async (form) => {
                        onFileHandle({ type: 'file-image-convert', config: values })
                        hide()
                    }
                "
                >{{ t('btn.submit') }}</FormButton
            >
        </div>
    </VeeForm>
</template>
