package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Task struct {
	ID uuid.UUID		`gorm:"type:uuid;column:id;default:gen_random_uuid()" json:"id"`
	Title string		 `gorm:"column:title" json:"title"`
	Description string	 `gorm:"column:description" json:"description"`
	FileName string		 `gorm:"column:fileName" json:"filename"`
	Subtask []Subtask    `gorm:"ForeignKey:Task_ID;association_foreignkey:ID" json:"subtask"`
	Status int			  `gorm:"column:status" json:"status"`
	CreatedAt time.Time	 `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
}

type AddTask struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	FileName string `json:"filename" binding:"required"`
	Status int `json:"status"`
}

type EditTask struct {
	Title string `json:"title"`
	Description string `json:"description"`
	FileName string `json:"filename"`
	Status int `json:"status"`
}


type Subtask struct {
	ID uuid.UUID			`gorm:"type:uuid;column:id;default:gen_random_uuid()" json:"id"`
	Title  string			`gorm:"column:title" json:"title"` 
	Description string 		`gorm:"column:description" json:"description"`
	Task_ID		string		`gorm:"column:task_id" json:"task_id"`
	Status 		int			`gorm:"column:status" json:"status"`
	CreatedAt time.Time		`gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  	`gorm:"column:updated_at" json:"-"`
}

type AddSubtask struct {
	Title string 			`json:"title" binding:"required"`
	Description string		`json:"description" binding:"required"`
	Task_ID string			`json:"task_id" binding:"required"`
	Status int 				`json:"status" binding:"required"`
}

type EditSubtask struct {
	Title string 			`json:"title" `
	Description string		`json:"description" `
	Task_ID string			`json:"task_id" `
	Status int 				`json:"status" `
}

type ErrorMessage struct {
	Message string
}
