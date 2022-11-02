package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"todoList/models"
	"todoList/pkg/setting"

	"github.com/labstack/echo/v4"
)


func GetTasks(e echo.Context) error {
	var perpageNum int = 100
	var err error
	var pageNum int
	var tasks []models.Task
	var offset int

	page := e.QueryParam("page")
	perpage := e.QueryParam("perpage")

	if len(page) != 0 && len(perpage) != 0 {
		if pageNum, err = strconv.Atoi(page); err != nil  {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		}
		if perpageNum , err = strconv.Atoi(perpage) ; err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		}
		offset = (pageNum - 1) * perpageNum

	}
	tasks, err = models.GetTasks(offset, perpageNum, false); 
	if err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		
	}
	return e.JSON(http.StatusOK, tasks)	
}

func GetTasksWithSubtask(e echo.Context) error {
	var tasks []models.Task
	var err error

	if tasks,err = models.GetTasks(0,0,true) ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, tasks)
}

func GetTask(e echo.Context) error {
	var task *models.Task
	var err error
	id := e.Param("id")
	if task,err = models.GetTask(id) ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	}
	return e.JSON(http.StatusOK, task)
}

func CreateTask(e echo.Context) error {
	var err error
	var newTask models.AddTask
	var dst *os.File

	file ,err := e.FormFile("file") ; 
	if err == nil {
		src, err := file.Open()
		defer src.Close()
	
		// Destination
		dst, err = os.Create(fmt.Sprintf("%s/%s", setting.FileHandlerSetting.Filepath, file.Filename))
		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	
		}
		defer dst.Close()
	
		if _, err = io.Copy(dst, src); err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		}
		newTask.FileName = file.Filename

	}else if err.Error() == "http: no such file" {
		//  JIKA TIDAK KIRIM FILE
	}else {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	}



	if err = e.Bind(&newTask) ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	}


	if err = models.CreateNewTask(newTask); err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		 
	}

	return e.JSON(http.StatusOK, newTask)
}

func UpdateTask(e echo.Context) error{
	var err error
	var updateTask models.EditTask
	var task *models.Task
	var dst *os.File
	id := e.Param("id")


	file,err := e.FormFile("file")

	if err == nil {
		src, err := file.Open()
		if err!= nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()}) 
		}

		defer src.Close()
		dst, err = os.Create(fmt.Sprintf("%s/%s", setting.FileHandlerSetting.Filepath, file.Filename))

		if err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()}) 
		}

		defer dst.Close()
		if _, err = io.Copy(dst,src) ; err != nil {
			return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()}) 

		}

		updateTask.FileName = file.Filename
	}else if err.Error() == "http: no such file"{

	}else {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()}) 
	}

	if err = e.Bind(&updateTask) ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
	}

	if task,err = models.UpdateTask(id, updateTask) ;err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		
	}
	return e.JSON(http.StatusOK, task)

}

func DeleteTask (e echo.Context) error{
	var err error
	var id string
	id = e.Param("id")
	if err = models.DeleteTask(id)  ; err != nil {
		return e.JSON(http.StatusBadRequest, models.ErrorMessage{Message: err.Error()})
		
	}
	
	return e.JSON(http.StatusOK, models.ErrorMessage{Message: "succes delete"})
}