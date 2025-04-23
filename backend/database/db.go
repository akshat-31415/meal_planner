package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	// dsn := GetDSN()  // Fetch connection string from config
	var err error
	err = godotenv.Load()
	if err != nil {
	    log.Fatal("Error loading .env file")
	}

	// Get environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Formulate the Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to the database")
}
