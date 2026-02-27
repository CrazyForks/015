package services

import (
	"errors"
	"os"
	"path/filepath"
	"pkg/models"
	"time"
	"worker/internal/utils"
)

type GenStandardFileReturn struct {
	FileId string
	models.FileInfo
}

// 生成标准格式的file
func GenStandardFile(filePath string, mimeType string) (GenStandardFileReturn, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return GenStandardFileReturn{}, errors.New("文件不存在")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return GenStandardFileReturn{}, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return GenStandardFileReturn{}, err
	}
	fileSize := fileInfo.Size()

	fileHash, err := utils.GetFileMd5(file)
	if err != nil {
		return GenStandardFileReturn{}, err
	}

	fileId := utils.GetFileId(fileHash, fileSize)

	uploadPath, err := utils.GetUploadDirPath()
	if err != nil {
		return GenStandardFileReturn{}, err
	}
	newPath := filepath.Join(uploadPath, fileId)
	if err := os.Rename(filePath, newPath); err != nil {
		return GenStandardFileReturn{}, err
	}
	models.SetRedisFileInfo(fileId, models.RedisFileInfo{
		FileInfo: models.FileInfo{
			FileSize: fileSize,
			FileHash: fileHash,
			MimeType: mimeType,
		},
		FileType:  models.FileTypeUpload,
		CreatedAt: time.Now().Unix(),
	})

	return GenStandardFileReturn{
		FileId: fileId,
		FileInfo: models.FileInfo{
			FileSize: fileSize,
			FileHash: fileHash,
			MimeType: mimeType,
		},
	}, nil
}
