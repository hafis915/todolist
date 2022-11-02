package models

import (
	"errors"
	"time"
)


func CreateNewTask(task AddTask ) (error) {
	var err error

	newTask := Task {
		Title : task.Title,
		Description: task.Description,
		FileName: task.FileName,
		Status: 0,
	}
	if err = db.Create(&newTask).Error ; err != nil  {
		return err
	}
	return nil
}

func GetTasks(page int, perpage int, include bool) ([]Task, error) {
 
	var tasks []Task
	var err error
	if include  {
		if err = db.Preload("Subtask").Find(&tasks).Error; err != nil {
			return nil, err
		}
	}else {
		err = db.Offset(page).Limit(perpage).Find(&tasks).Error
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	var task Task
	var err error

	if err = db.Preload("Subtask").Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(id string ,updatedData EditTask ) (*Task, error) {
	var task Task
	var err error

	if err = db.Where("id = ? ", id).Find(&task).Error ; err != nil {
		return nil, err
	}

	var newVal Task = Task{
		Title: updatedData.Title ,
		Description: updatedData.Description,
		FileName: updatedData.FileName,
		Status: updatedData.Status,
		UpdatedAt: time.Now(),
	}

	if err = db.Model(&Task{}).Where("id = ?", id).Updates(newVal).Error; err != nil {
		return nil, err
	}

	if err = db.Where("id = ?", id).First(&task).Error ; err != nil {
		return nil, err
	}

	return &task, nil
}

func DeleteTask(id string) error {
	var task Task 
	var err error
	if err = db.Where("id= ?", id).Find(&task).Error ; err != nil {
		return errors.New("Data Not Found")
	}

	if err = db.Delete(&task).Error ; err != nil {
		return err
	}

	return nil
}

