package main

import (
	"github.com/daisuke8000/server/src/routes"
	"github.com/gin-contrib/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		//AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"},
		//AllowHeaders:     []string{"Origin"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc:  nil,
		//MaxAge:           12 * time.Hour,
	}))
	routes.Setup(app)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	app.Run()
}
