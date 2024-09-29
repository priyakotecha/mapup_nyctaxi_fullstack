package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var dbConfig *DBConfig

// Initialize the environment variables from the .env file
func LoadConfig() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Parse the DB_PORT environment variable as an integer
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT value: %v", err)
	}

	dbConfig = &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

// Getter for DBConfig
func GetDBConfig() *DBConfig {
	if dbConfig == nil {
		log.Fatal("Config not loaded. Call LoadConfig() before GetDBConfig().")
	}
	return dbConfig
}

// Build PostgreSQL connection string
func GetPostgresConnectionString() string {
	cfg := GetDBConfig()
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
}
