package routes

import (
	"github.com/daisuke8000/server/src/controllers"
	"github.com/daisuke8000/server/src/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	api := app.Group("/api")
	admin := api.Group("/admin")
	{
		admin.POST("/signup", controllers.Signup)
		admin.POST("/signin", controllers.Signin)
		adminAuthorization := admin.Use(middleware.IsAuthorization)
		adminAuthorization.GET("/user", controllers.User)
		adminAuthorization.POST("/signout", controllers.Signout)
		adminAuthorization.PUT("/users/update", controllers.UpdateInfo)
		adminAuthorization.PUT("/users/password", controllers.UpdatePassword)
		adminAuthorization.GET("/ambassadors", controllers.Ambassadors)
	}
	ambassador := api.Group("/ambassador")
	{
		ambassador.POST("/signup", controllers.Signup)
		ambassador.POST("/signin", controllers.Signin)
		ambassadorAuthorization := ambassador.Use(middleware.IsAuthorization)
		ambassadorAuthorization.GET("/user", controllers.User)
		ambassadorAuthorization.POST("/signout", controllers.Signout)
		ambassadorAuthorization.GET("/histories", controllers.Histories)
		ambassadorAuthorization.POST("/history", controllers.CreateHistory)
		ambassadorAuthorization.GET("/history/:id", controllers.GetHistory)
		ambassadorAuthorization.PUT("/history/:id", controllers.UpdateHistory)
		ambassadorAuthorization.PUT("/users/update", controllers.UpdateInfo)
		ambassadorAuthorization.PUT("/users/password", controllers.UpdatePassword)
		ambassadorAuthorization.GET("/users/point")
		ambassadorAuthorization.GET("/users/score")
	}
}
