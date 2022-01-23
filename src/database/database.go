package database

import (
	"github.com/daisuke8000/server/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/develop"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Match{}, models.History{})
}
