<script setup lang="ts">
import { useQueries, useQuery } from '@tanstack/vue-query'
import { AsyncButton, Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { filesize } from 'filesize'
import useMyAppShare from '@/composables/useMyAppShare'
import { toast } from 'vue-sonner'
import type { handleFileComponentProps } from './types'
const emit = defineEmits<{
    (e: 'change', key: string): void
}>()
const props = defineProps<handleFileComponentProps>()
const fileIds = computed(() => props?.data?.files?.map((item) => item.id))
const { data: taskIds } = useQuery({
    queryKey: ['create-image-compress', fileIds.value],
    queryFn: async () => {
        return await Promise.all(
            props?.data?.files?.map(async (file) => {
                const { id } = file || {}
                if (!id) return
                const data = await $fetch<{
                    code: number
                    data: {
                        id?: string
                    }
                }>(`/api/task/image:compress`, {
                    method: 'POST',
                    body: {
                        file_id: id,
                    },
                })
                return data?.data?.id
            })
        )
    },
    staleTime: Infinity,
    enabled: !!fileIds.value && fileIds.value?.length > 0,
})

const taskResults = useQueries({
    queries:
        taskIds?.value?.filter(Boolean).map((taskId) => {
            return {
                queryKey: ['task-image-compress', taskId],
                queryFn: async () => {
                    const data = await $fetch<{
                        code: number
                        data: {
                            result: {
                                old_file: {
                                    id: string
                                    size: number
                                }
                                new_file: {
                                    id: string
                                    size: number
                                }
                            }[]
                            status: 'success' | 'retry' | 'archived'
                            err?: {
                                message?: string
                                retry?: number
                                max_retry?: number
                            }
                        }
                    }>(`/api/task/${taskId}`)
                    return data?.data
                },
                enabled: !!taskId,
            }
        }) ?? [],
})

const { downloadFileByShareId, createFileShare } = useMyAppShare()

const { counter, pause } = useInterval(2000, { controls: true })

watch(
    () => counter.value,
    () => {
        taskResults.value.forEach((item) => {
            if (['success', 'archived'].includes(item.data?.status ?? '')) {
                pause()
                return
            }
            item.refetch()
        })
    }
)

watch(
    () => taskResults.value,
    (newVal, oldVal) => {
        newVal?.map((item, index) => {
            const { retry, max_retry, message } = item.data?.err || {}
            if (!retry || !max_retry || retry >= max_retry || retry === oldVal?.[index]?.data?.err?.retry) return
            toast.error(`处理错误: ${message}, 将再次重试`)
        })
    }
)
</script>
<template>
    <div class="flex flex-col gap-3">
        <h2 class="text-lg">上传成功</h2>
        <div class="flex flex-col gap-1 items-center">
            <div class="flex flex-col h-30 items-center justify-center">
                <FilePreviewView :value="props?.data?.file" />
            </div>
        </div>
        <div v-if="taskData?.status === 'success'" class="flex flex-col gap-2" v-for="item in taskData?.result">
            <div class="bg-white/80 p-2 rounded-md w-full flex flex-row items-center justify-between gap-2">
                <div class="flex flex-row gap-2 items-center max-w-2/3">
                    <div class="flex flex-row items-center justify-center rounded-md bg-black/5 p-2">
                        <LucideImage />
                    </div>
                    <div class="truncate w-auto">{{ props?.data?.file?.name }}</div>
                    <div class="flex flex-row gap-2 items-center text-sm shrink-0">
                        <span class="opacity-75">{{ filesize(item.new_file.size ?? 0) }}</span>
                        <span class="bg-green-200 text-green-600 rounded-md px-1 py-0.5 flex flex-row gap-1 items-center text-xs">
                            <LucideChevronDown class="size-4" />
                            {{ ((1 - item.new_file.size / item.old_file.size) * 100).toFixed(2) }}%
                        </span>
                    </div>
                </div>
                <AsyncButton
                    variant="outline"
                    class="bg-black/5"
                    size="icon"
                    @click="
                        async () => {
                            const data = await createFileShare({
                                file_id: item.new_file.id,
                                config: {
                                    download_nums: 1,
                                    expire_time: 60,
                                    has_pickup_code: false,
                                    has_password: false,
                                },
                                file_name: props?.data?.file?.name,
                            })
                            const { id } = data?.data || {}
                            if (!id) {
                                return
                            }
                            try {
                                await downloadFileByShareId(id)
                            } catch (error: any) {
                                toast.error(error?.data?.message || error?.message || error)
                            }
                        }
                    "
                >
                    <LucideDownload />
                </AsyncButton>
            </div>
        </div>
        <div v-else-if="taskData?.status !== 'retry' && !!taskData?.err?.message" class="flex flex-col gap-2">
            <div class="w-full h-16 flex flex-row items-center gap-3 bg-white/80 rounded-md p-2">
                <div class="size-10 flex items-center justify-center rounded-md bg-red-200">
                    <LucideAlertTriangle class="size-5 text-red-600" />
                </div>
                <div class="text-sm">
                    {{ `经过 ${taskData?.err?.retry} 次重试后任务处理失败: ${taskData?.err?.message}` }}
                </div>
            </div>
            <div class="flex flex-row justify-center">
                <Button @click="emit('change', 'input')"> 返回首页 </Button>
            </div>
        </div>

        <div v-else="taskData?.status !== 'retry' && !!taskData?.err?.message" class="flex flex-col gap-2">
            <Skeleton class="w-full h-16 flex flex-row items-center justify-between" v-for="i in 3" />
        </div>
    </div>
</template>
