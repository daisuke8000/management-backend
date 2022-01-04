package main

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.AutoMigrate()
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
