package models

import (
	"errors"
	"time"
)


func CreateSubtask(task *AddSubtask) error {
	var err error
	newSubTask := Subtask{
		Title: task.Title,
		Description: task.Description,
		Status: 0,
		Task_ID: task.Task_ID,
	}
	if err = db.Create(&newSubTask).Error; err != nil {
		return err
	}

	return nil
}

func GetSubtasks() ([]Subtask, error) {
	var err error
	var subtask []Subtask
	if err = db.Find(&subtask).Error; err != nil {
		return nil, err
	}
	return subtask, nil
}

func GetSubtask(task_id string) (*[]Subtask, error) {
	var subtask []Subtask
	var err error
	if err = db.Where("task_id= ? ", task_id).Find(&subtask).Error ; err != nil {
		return nil, err
	}
	return &subtask, nil
}

func UpdateSubtask(id string, updatedData EditSubtask) (*Subtask, error) {
	var subtask Subtask
	var err error
	if err = db.Where("id = ? ", id).First(&subtask).Error; err != nil {
		return nil, errors.New("Data Not Found")
	}

	var newVal Subtask = Subtask {
		Title: updatedData.Title,
		Description :updatedData.Description,
		Task_ID: updatedData.Task_ID,
		Status: updatedData.Status,
		UpdatedAt: time.Now(),
	}
	if err = db.Model(&Subtask{}).Where("id=?", id).Updates(newVal).Error; err != nil {
		return nil, err
	}

	if err = db.Where("id= ?", id).Find(&subtask).Error ; err != nil {
		return nil, err
	}

	return &subtask, err
}

func DeleteSubtask(id string) error {
	var err error
	var subtask Subtask
	if err = db.Where("id= ?", id).First(&subtask).Error; err != nil {
		return err
	}

	if err = db.Delete(&subtask).Error; err != nil  {
		return err
	}
	return nil
}