package main

import (
	"github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	"mealplan/database"
	"mealplan/routes"
)

func main() {
	database.Connect()

	database.Migrate()

	r := routes.SetupRoutes()

	r.Use(cors.New(cors.Config{
    	AllowOrigins: []string{"https://https://meal-planner-rosy.vercel.app/"},
    	AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}))

	r.Run("0.0.0.0:8080")

}
