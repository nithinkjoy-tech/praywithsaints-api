package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Retrieve database URL from environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatalf("DATABASE_URL environment variable not set")
	}

	// Open a database connection
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	// Verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Unable to verify connection: %v", err)
	}

	fmt.Println("Database connected!")
}

func InitializePrayersTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS prayers (
            prayer_id SERIAL PRIMARY KEY,
            prayer_name VARCHAR(255) NOT NULL,
            how_to_pray TEXT,
            impact TEXT,
            talk_link VARCHAR(255),
            author VARCHAR(255),
            testimonies TEXT[],
            references TEXT[]
        )
    `
	_, err := DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
