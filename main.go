package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	var err error
	//database.Connect()
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/develop"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}
	//database.AutoMigrate()

	app := gin.Default()

	//routes.Setup(app)

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	app.Run()
}
