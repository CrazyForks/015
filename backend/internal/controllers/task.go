package controllers

import (
	"backend/internal/controllers/task"
	"backend/internal/utils"
	"backend/middleware"
	"errors"
	"pkg/models"
	u "pkg/utils"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
)

var handleTaskMap = map[string]func(c *middleware.CustomContext) ([]byte, error){
	"image:compress": task.HandleImageCompress,
}

func CreateTask(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	taskType := cc.Param("type")
	if taskType == "" {
		return utils.HTTPErrorHandler(c, errors.New("调用接口参数错误"))
	}
	handleTask, ok := handleTaskMap[taskType]
	if !ok {
		return utils.HTTPErrorHandler(c, errors.New("任务不存在"))
	}
	json, err := handleTask(cc)
	if err != nil {
		return utils.HTTPErrorHandler(c, err)
	}

	client := u.GetQueueClient()
	info, err := client.Enqueue(asynq.NewTask(taskType, json), asynq.MaxRetry(3))
	if err != nil {
		return utils.HTTPErrorHandler(c, err)
	}

	return utils.HTTPSuccessHandler(c, map[string]any{
		"id": info.ID,
	})
}

func GetTask(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	taskId := cc.Param("id")
	if taskId == "" {
		return utils.HTTPErrorHandler(c, errors.New("调用接口参数错误"))
	}

	taskInfo, err := models.GetRedisTaskInfo(taskId)
	if err != nil {
		return utils.HTTPErrorHandler(c, err)
	}
	if taskInfo == nil {
		client := u.GetQueueInspector()

		queneTaskInfo, err := client.GetTaskInfo("default", taskId)
		if err != nil {
			return utils.HTTPErrorHandler(c, errors.New("任务已过期"))
		}
		stateMap := map[asynq.TaskState]string{
			asynq.TaskStateActive:    "processing",
			asynq.TaskStatePending:   "pending",
			asynq.TaskStateScheduled: "scheduled",
			asynq.TaskStateRetry:     "retry",
			asynq.TaskStateArchived:  "archived",
			asynq.TaskStateCompleted: "completed",
		}
		if queneTaskInfo != nil {
			return utils.HTTPSuccessHandler(c, map[string]any{
				"status": stateMap[queneTaskInfo.State],
				"err": map[string]any{
					"message":   queneTaskInfo.LastErr,
					"retry":     queneTaskInfo.Retried,
					"max_retry": queneTaskInfo.MaxRetry,
				},
			})
		}
	}
	return utils.HTTPSuccessHandler(c, *taskInfo)
}
