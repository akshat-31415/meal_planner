package main

import (
	"github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	"mealplan/database"
	"mealplan/routes"
	"os"
	"fmt"
)

func main() {
	database.Connect()

	database.Migrate()

	r := routes.SetupRoutes()

	r.Use(cors.New(cors.Config{
    	AllowOrigins: []string{"https://https://meal-planner-rosy.vercel.app"},
    	AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}))

	//r.Run("0.0.0.0:8080")
	port := os.Getenv("PORT")
    	if port == "" {
	        port = "10000" // Render uses 10000 by default if PORT is not set
	    }
	
	fmt.Printf("Starting server on port %s\n", port)  // Log the port for debugging
	r.Run("0.0.0.0:" + port) // Bind to all interfaces on the specified port


}
