package database

import (
	"log"
	"mealplan/models"
)

// Migrate performs database migrations, creating tables or altering schemas
func Migrate() {
	// Automatically migrate your models
	err := DB.AutoMigrate(&models.User{}/*, &models.Meal{}, &models.FoodItem{}*/)
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
}
