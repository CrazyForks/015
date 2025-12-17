import type { FileHandleKey, TextHandleKey } from '../Preprocessing/types'

type basehandleData = { config: Record<string, any> }

export type filehandleData = { files: { id: string; file: File }[]; handle_type: FileHandleKey } & basehandleData
export type texthandleData = { text: string; handle_type: TextHandleKey } & basehandleData
export type handleKey = FileHandleKey | TextHandleKey

export type handleTextComponentProps = { data: texthandleData }
export type handleFileComponentProps = { data: filehandleData }
export type handleComponent = Component<handleTextComponentProps | handleFileComponentProps>
