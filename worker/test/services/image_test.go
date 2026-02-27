package services

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"worker/internal/services"
	"worker/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestCompressPNGHappyPath(t *testing.T) {
	tmp := t.TempDir()
	filePath := filepath.Join(tmp, "test.png")
	// 从 test/resource 复制真实 PNG，路径基于当前测试文件位置
	_, self, _, _ := runtime.Caller(0)
	srcPath := filepath.Join(filepath.Dir(self), "..", "resource", "test.png")
	err := utils.CopyFile(srcPath, filePath)
	if err != nil {
		t.Fatal(err)
	}
	got, err := services.CompressImage(filePath, "image/png")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, filePath+"_compressed")
	origInfo, err := os.Stat(filePath)
	if err != nil {
		t.Fatal(err)
	}
	compInfo, err := os.Stat(got)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, origInfo.Size(), compInfo.Size())
	fmt.Printf("原图: %d | 压缩后: %d | 压缩率: %f%%\n", origInfo.Size(), compInfo.Size(), float64(compInfo.Size())/float64(origInfo.Size())*100)
}

func TestCompressJPEGHappyPath(t *testing.T) {
	tmp := t.TempDir()
	filePath := filepath.Join(tmp, "test.jpg")
	_, self, _, _ := runtime.Caller(0)
	srcPath := filepath.Join(filepath.Dir(self), "..", "resource", "test.jpg")
	err := utils.CopyFile(srcPath, filePath)
	if err != nil {
		t.Fatal(err)
	}
	got, err := services.CompressImage(filePath, "image/jpeg")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got, filePath+"_compressed")
	origInfo, err := os.Stat(filePath)
	if err != nil {
		t.Fatal(err)
	}
	compInfo, err := os.Stat(got)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, origInfo.Size(), compInfo.Size())
	fmt.Printf("原图: %d | 压缩后: %d | 压缩率: %f%%\n", origInfo.Size(), compInfo.Size(), float64(compInfo.Size())/float64(origInfo.Size())*100)
}

func TestConvertImageWithMagickPNGToJPG(t *testing.T) {
	tmp := t.TempDir()
	filePath := filepath.Join(tmp, "test.png")
	// 从 test/resource 复制真实 PNG
	_, self, _, _ := runtime.Caller(0)
	srcPath := filepath.Join(filepath.Dir(self), "..", "resource", "test.png")
	err := utils.CopyFile(srcPath, filePath)
	if err != nil {
		t.Fatal(err)
	}

	// 测试 PNG 转 JPEG
	got, err := services.ConvertImageWithMagick(filePath, "image/png", ".jpg")
	if err != nil {
		t.Fatal(err)
	}

	// 验证输出文件路径
	assert.Equal(t, got, filePath+"_converted.jpg")

	// 验证输出文件存在
	info, err := os.Stat(got)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, info.Size() > 0, "转换后的文件应该有内容")

	fmt.Printf("PNG 转 JPEG 成功: %s (大小: %d bytes)\n", got, info.Size())
}

func TestConvertImageWithMagickInvalidExt(t *testing.T) {
	tmp := t.TempDir()
	filePath := filepath.Join(tmp, "test.png")
	_, self, _, _ := runtime.Caller(0)
	srcPath := filepath.Join(filepath.Dir(self), "..", "resource", "test.png")
	err := utils.CopyFile(srcPath, filePath)
	if err != nil {
		t.Fatal(err)
	}

	// 测试非法扩展名（防止注入）
	_, err = services.ConvertImageWithMagick(filePath, "image/png", ".exe")
	assert.Error(t, err, "应该返回错误")
	assert.Equal(t, services.ErrInvalidImageExt, err, "应该返回 ErrInvalidImageExt 错误")
}
