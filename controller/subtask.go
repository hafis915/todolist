package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"todoList/models"
	"todoList/pkg/setting"

	"github.com/labstack/echo/v4"
)


func GetSubtasks(e echo.Context) error {
	subtasks, err := models.GetSubtasks(); 
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		
	}
	return e.JSON(http.StatusOK, subtasks)	
}

func GetSubtask(e echo.Context) error {
	var subtask *[]models.Subtask
	var err error
	taskid := e.Param("taskid")
	if subtask,err = models.GetSubtask(taskid) ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	}
	return e.JSON(http.StatusOK, subtask)
}

func CreateSubTask(e echo.Context) error {
	var err error
	var newSubtask models.AddSubtask
	var dst *os.File

	file, err := e.FormFile("file")
	
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()}) 
		}
		defer src.Close()

		dst, err = os.Create(fmt.Sprintf("%s/%s", setting.FileHandlerSetting.Filepath, file.Filename))
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()}) 
		}

		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		}

		newSubtask.FileName = file.Filename

	}else if err.Error() == "http: no such file" {

	}else {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	}

	if err = e.Bind(&newSubtask) ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()}) 
		
	}

	if err = models.CreateSubtask(&newSubtask); err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		 
	}

	return e.JSON(http.StatusOK, newSubtask)
}

func UpdateSubTask(e echo.Context) error{
	var err error
	var updateSubtask models.EditSubtask
	var subtask *models.Subtask
	id := e.Param("id")
	if err = e.Bind(&updateSubtask) ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		
	}

	if subtask,err = models.UpdateSubtask(id, updateSubtask) ;err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		
	}

	return e.JSON(http.StatusOK, subtask)

}

func DeleteSubtask (e echo.Context) error{
	var err error
	var id string
	id = e.Param("id")
	if err = models.DeleteSubtask(id)  ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		
	}
	
	return e.JSON(http.StatusOK, models.ErrorMessage{Message: "succes delete"})
}