package routes

import "github.com/gin-gonic/gin"

func Setup(app *gin.Engine) {
	api := app.Group("/api")
	admin := api.Group("/admin")
	{
		admin.GET("/signup", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "admin signup",
			})
		})
	}
	ambassador := api.Group("/ambassador")
	{
		ambassador.GET("/signup", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ambassador signup6",
			})
		})
	}
}
