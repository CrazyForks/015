package tasks

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"pkg/models"
	u "pkg/utils"
	"worker/internal/utils"

	"github.com/hibiken/asynq"
)

func RemoveFile(ctx context.Context, task *asynq.Task) error {
	var payload RemoveFileTaskPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	fileInfo, err := models.GetRedisFileInfo(payload.FileId)
	if err != nil {
		return err
	}
	if fileInfo == nil || fileInfo.FileType == models.FileTypeUpload {
		return nil
	}

	rdb, rctx := u.GetRedisClient()
	uploadPath, err := utils.GetUploadDirPath()
	if err != nil {
		return err
	}
	filePath := filepath.Join(uploadPath, payload.FileId)
	rdb.HDel(rctx, "015:fileInfoMap", payload.FileId)
	os.RemoveAll(filePath)
	return nil
}
