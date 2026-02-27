package services

import (
	"errors"
	"worker/internal/utils"

	"github.com/samber/lo"
)

var (
	ErrUnsupportedMimeType = errors.New("UnsupportedMimeType")
)

// 支持的图片扩展名列表
var supportedImageExts = []string{
	".jpg", ".jpeg", ".png", ".gif", ".webp",
	// ".bmp", ".tiff", ".tif", ".svg", ".ico",
}

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
		return "", ErrUnsupportedMimeType
	}
	return compressedPath, nil
}

func ConvertImageWithMagick(filePath, mimeType, targetExt string) (string, error) {
	// 验证目标扩展名是否合法
	if !lo.Contains(supportedImageExts, targetExt) {
		return "", ErrUnsupportedMimeType
	}

	outputPath := filePath + "_converted" + targetExt

	_, err := utils.RunCommand("magick", filePath, outputPath)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
