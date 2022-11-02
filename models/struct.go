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
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	FileName string 
}

type EditTask struct {
	Title string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	FileName string
	Status int `json:"status" form:"status"`
}


type Subtask struct {
	ID uuid.UUID			`gorm:"type:uuid;column:id;default:gen_random_uuid()" json:"id"`
	Title  string			`gorm:"column:title" json:"title"` 
	Description string 		`gorm:"column:description" json:"description"`
	Task_ID		string		`gorm:"column:task_id" json:"task_id"`
	Status 		int			`gorm:"column:status" json:"status"`
	FileName	string    	`gorm:"column:fileName" json:"filename"`
	CreatedAt time.Time		`gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  	`gorm:"column:updated_at" json:"-"`
}

type AddSubtask struct {
	Title string 			`json:"title" form:"title" binding:"required"`
	Description string		`json:"description" form:"description" binding:"required"`
	Task_ID string			`json:"task_id" form:"task_id" binding:"required"`
	FileName string
} 

type EditSubtask struct {
	Title string 			`json:"title" form:"title"`
	Description string		`json:"description" form:"description"`
	Task_ID string			`json:"task_id" form:"task_id"`
	Status int 				`json:"status" form:"status"`
	FileName string
}

type ErrorMessage struct {
	Message string
}
