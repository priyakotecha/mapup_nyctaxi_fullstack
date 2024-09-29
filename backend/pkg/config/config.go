package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	host             string
	port             int
	user             string
	password         string
	dbName           string
	redisURL         string
	paraquetURL      string
	paraquetFilePath string
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
		host:             os.Getenv("DB_host"),
		port:             port,
		user:             os.Getenv("DB_USER"),
		password:         os.Getenv("DB_PASSWORD"),
		dbName:           os.Getenv("DB_NAME"),
		redisURL:         os.Getenv("REDIS_URL"),
		paraquetURL:      os.Getenv("NYC_TRIP_DATA_PARAQUET_FILE_URL"),
		paraquetFilePath: "/data/trip_data",
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
		cfg.host, cfg.port, cfg.user, cfg.password, cfg.dbName)

}

func GetParaquetFileURLAndPath() (string, string) {
	cfg := GetDBConfig()
	return cfg.paraquetURL, cfg.paraquetFilePath
}
