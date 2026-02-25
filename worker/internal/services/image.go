package services

import (
	"errors"
	"worker/internal/utils"
)

func CompressImage(filePath string, mimeType string) (string, error) {
	compressedPath := filePath + "_compressed"
	switch mimeType {
	case "image/png":
		_, err := utils.RunCommand("pngquant", "--output", compressedPath, filePath)
		if err != nil {
			return "", err
		}
	case "image/jpeg":
		err := utils.CopyFile(filePath, compressedPath)
		if err != nil {
			return "", err
		}
		_, err = utils.RunCommand("jpegoptim", "-m", "80", "--strip-all", compressedPath)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("不支持的文件类型")
	}
	return compressedPath, nil
}
