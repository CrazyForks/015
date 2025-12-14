package main

import (
	"fmt"
	"pkg/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	// 日志
	var logger *zap.Logger
	if utils.GetEnvWithDefault("node.env", "production") == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	e := echo.New()
	for _, middleware := range middlewares {
		e.Use(middleware())
	}

	for _, route := range routes {
		e.Match(route.Method, route.Path, route.Handler)
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", utils.GetEnvWithDefault("api.port", "5001"))))
}
