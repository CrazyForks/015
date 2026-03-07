type Uploadfile = {
    fileId: string
    id?: string // 后端文件id
    file: File
    status: 'start' | 'pause' | 'finish' | 'error'
    hash?: string
    procressType: 'hash' | 'create' | 'upload' | 'finish'
    uploadInfo?: {
        chunks: Record<number, { status: 'success' | 'error' | 'processing'; createdAt: number }>
        chunkLength: number
        ChunkSize: number
    }
    queue: {
        taskId: string
        taskType: 'hash' | 'create' | 'chunk' | 'upload' | 'finish'
        queueType: 'sync' | 'async' // sync任务禁止并发
        index?: number
        retry?: number
    }[]
}

type SelectedFile = string | null

export type { Uploadfile, SelectedFile }
