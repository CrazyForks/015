package task

import (
	"encoding/json"
	"errors"

	"github.com/labstack/echo/v5"
)

type GenCompressImageRequest struct {
	FileId string `json:"file_id"`
}

func HandleImageCompress(c *echo.Context) ([]byte, error) {

	r := new(GenCompressImageRequest)
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
