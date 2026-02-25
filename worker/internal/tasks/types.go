package tasks

type BaseFileTaskPayload struct {
	FileId string `json:"file_id"`
}

type RemoveFileTaskPayload struct {
	BaseFileTaskPayload
}

type CompressImageTaskPayload struct {
	BaseFileTaskPayload
}
