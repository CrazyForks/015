package task

import (
	"backend/middleware"
	"encoding/json"
	"errors"
)

type GenCompressImageRequest struct {
	FileId string `json:"file_id"`
}

func HandleImageCompress(c *middleware.CustomContext) ([]byte, error) {

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
