package routes

import (
	"mealplan/controllers"
	"mealplan/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.GET("/:id", middleware.AuthMiddleware(), controllers.Profile)
	}
	// r.Static("/", "./frontend")
	r.GET("/", func(c *gin.Context) {
		c.File("../frontend/index.html")
	})

	return r
}
