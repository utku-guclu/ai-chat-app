package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	/*
	*In Go, if you import a package just for its side effects (like the PostgreSQL driver, which registers itself to the database/sql package), you must prefix the import path with an underscore (_).
	* */
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global database connection pool
var DB *sql.DB

// InitSchema creates the necessary tables in the database if they do not exist.
func InitSchema() error {
	// Users Table
	// The username column is UNIQUE and NOT NULL to prevent duplicates
	_, err := DB.Exec(`
            CREATE TABLE IF NOT EXISTS users (
                id SERIAL PRIMARY KEY,
                username VARCHAR(255) UNIQUE NOT NULL,
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
            );
        `)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Messages Table
	// user_id links to the users table via a foreign key
	_, err = DB.Exec(`
            CREATE TABLE IF NOT EXISTS messages (
                id SERIAL PRIMARY KEY,
                user_id INTEGER NOT NULL REFERENCES users(id),
                content TEXT NOT NULL,
                is_ai BOOLEAN DEFAULT FALSE,
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
            );
        `)
	if err != nil {
		return fmt.Errorf("failed to create messages table: %w", err)
	}

	fmt.Println("Database schema initialized successfully.")
	return nil
}

// ConnectDB initializes the PostgreSQL connection using environment variables.
func ConnectDB() {
	// 1. Get connection details from environment variables (set in docker-compose.yml)
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	// 2. Create the connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, db_name)

	var err error
	// 3. Open the connection
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		// We use log.Fatal here, which will stop the container if connection fails
		log.Fatal("Error opening database connection", err)
	}

	// 4. Set pool limits (optional, but good practice for performance)
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// 5. Verify the connection (ping)
	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
}
