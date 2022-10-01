package routes

import (
	"api/controllers"
	"api/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}

		userauth := main.Group("user", middlewares.Auth())
		{
			userauth.GET("/list", controllers.ListUsers)
			user.GET("/my_profile", controllers.GetUser)
		}

		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
	}
	return router
}
