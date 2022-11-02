package router

import (
	"todoList/controller"

	"github.com/labstack/echo/v4"
)


func TaskRoute(echo *echo.Echo) {
	echo.GET("/tasksubtask", controller.GetTasksWithSubtask)
	echo.GET("/tasks", controller.GetTasks)
	echo.GET("/task/:id", controller.GetTask)
	echo.POST("/task", controller.CreateTask)
	echo.PUT("/task/:id", controller.UpdateTask)
	echo.DELETE("/task/:id", controller.DeleteTask)
} 

func SubtaskRoute (echo *echo.Echo) {
	echo.GET("/subtasks", controller.GetSubtasks)
	echo.GET("/subtask/:taskid", controller.GetSubtask)
	echo.POST("/subtask", controller.CreateSubTask)
	echo.PUT("/subtask/:id", controller.UpdateSubTask)
	echo.DELETE("/subtask/:id", controller.DeleteSubtask)
}