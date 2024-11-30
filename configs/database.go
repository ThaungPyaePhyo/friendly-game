package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {

	envPath := filepath.Join("configs", ".env")	
	if err := godotenv.Load(envPath); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    hostname := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	var err error

    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	fmt.Println("Connected to MySQL database successfully!")
}