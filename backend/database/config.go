package database

import (
	"os"
	"fmt"
)

// GetDSN returns the DSN (Data Source Name) for the PostgreSQL connection
func GetDSN() string {
	// Ideally, these should be environment variables or in a .env file
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
}
