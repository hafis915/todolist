package models

import (
	"fmt"
	"log"
	"todoList/pkg/setting"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)



var db *gorm.DB
func Setup()  {
	var err error
	fmt.Println(setting.DatabaseSetting.SSL)
	db, err = gorm.Open(setting.DatabaseSetting.Type,fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s", setting.DatabaseSetting.Type, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.DBName, setting.DatabaseSetting.SSL) )

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&Task{}, &Subtask{})
}