package main

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("user_jwt", store))
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3007"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{
			"Content-Type",
		},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc:  nil,s
		MaxAge: 12 * time.Hour,
	}))
	routes.Setup(app)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	app.Run()
}
