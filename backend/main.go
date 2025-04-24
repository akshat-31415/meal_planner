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

	r.Use(cors.Default())

	r.Run(":8080")

}
