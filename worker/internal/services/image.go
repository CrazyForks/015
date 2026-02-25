package services

import (
	"errors"
	"os/exec"
	"worker/internal/utils"
)

func CompressImage(filePath string, mimeType string) (string, error) {
	compressedPath := filePath + "_compressed"
	switch mimeType {
	case "image/png":
		args := []string{"--output", compressedPath, filePath}
		cmd := exec.Command("pngquant", args...)
		_, err := cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
	case "image/jpeg":
		err := utils.CopyFile(filePath, compressedPath)
		if err != nil {
			return "", err
		}
		args := []string{"-m", "90", "--strip-all", compressedPath}
		cmd := exec.Command("jpegoptim", args...)
		_, err = cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("不支持的文件类型")
	}
	return compressedPath, nil
}
