package routes

import (
	"github.com/daisuke8000/server/src/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	api := app.Group("/api")
	admin := api.Group("/admin")
	{
		admin.POST("/signup", controllers.Signup)
		admin.POST("/signin", controllers.Signin)
		//adminAuthorization := admin.Use(middleware.IsAuthorization)
	}

	ambassador := api.Group("/ambassador")
	{
		ambassador.GET("/signup", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ambassador signup8",
			})
		})
	}
}
