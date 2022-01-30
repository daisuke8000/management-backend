package main

import (
	"github.com/bxcodec/faker/v3"
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
)

func main() {
	database.Connect()
	for i := 0; i < 400; i++ {
		ambassador := models.User{
			Name:    faker.Name(),
			Email:   faker.Email(),
			IsAdmin: false,
		}
		ambassador.SetPassword("9999")
		database.DB.Create(&ambassador)
	}
}
