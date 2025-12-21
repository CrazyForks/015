package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"pkg/utils"

	humanize "github.com/dustin/go-humanize"
)

func GetFileId(fileHash string, fileSize int64) string {
	return fmt.Sprintf("%s_%d", fileHash, fileSize)
}

func GetFileMd5(file io.Reader) (string, error) {

	const bufferSize = 1024 * 1000 // 1MB

	hash := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		hash.Write(buf[:n])
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func GetUploadDirPath() (string, error) {
	uploadPath := utils.GetEnv("upload.path")
	if uploadPath == "" {
		basepath, err := os.Getwd()
		if err != nil {
			return "", err
		}
		uploadPath = filepath.Join(basepath, "uploads")
	}
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", err
	}
	return uploadPath, nil
}

func GetFileSize(size string) (uint64, error) {
	return humanize.ParseBytes(size)
}
