package task

import (
	"encoding/json"
	"errors"

	"github.com/labstack/echo/v5"
)

type BaseImageRequest struct {
	FileId string `json:"file_id"`
}

func HandleImageCompress(c *echo.Context) ([]byte, error) {
	r := new(BaseImageRequest)
	if err := c.Bind(r); err != nil {
		return nil, err
	}
	if r.FileId == "" {
		return nil, errors.New("调用接口参数错误")
	}
	json, err := json.Marshal(map[string]any{
		"file_id": r.FileId,
	})
	if err != nil {
		return nil, err
	}

	return json, nil
}

type ImageConvertRequest struct {
	BaseImageRequest
	TargetExt string `json:"target_ext"`
}

func HandleImageConvert(c *echo.Context) ([]byte, error) {
	r := new(ImageConvertRequest)
	if err := c.Bind(r); err != nil {
		return nil, err
	}
	if r.FileId == "" {
		return nil, errors.New("调用接口参数错误")
	}
	json, err := json.Marshal(map[string]any{
		"file_id": r.FileId,
	})
	if err != nil {
		return nil, err
	}

	return json, nil
}
