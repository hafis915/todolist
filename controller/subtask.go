package controller

import (
	"net/http"
	"todoList/models"
	"todoList/utility"

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

	file, err := e.FormFile("file")
	
	if err == nil {
		filename, err := utility.UploadFile(file)
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})

		}
		newSubtask.FileName = *filename

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

	file, err := e.FormFile("file")

	if err == nil {
		filename, err := utility.UploadFile(file)
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		}
		updateSubtask.FileName = *filename
	}else if err.Error() == "http: no such file" {

	}else {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})

	}

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