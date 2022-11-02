package main

import (
	"todoList/models"
	"todoList/pkg/setting"
	"todoList/router"

	"github.com/labstack/echo/v4"
)

func initialize() {
	setting.Setup()
	models.Setup()
}

func main() {
	initialize()
	e := echo.New()

	router.SubtaskRoute(e)
	router.TaskRoute(e)

	e.Start(setting.ServerSetting.HttpPort)
}