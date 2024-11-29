package database

import (
    "database/sql"
    "fmt"
	"log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	username := "tpp"            
	password := "password"       
	hostname := "127.0.0.1:3306" 
	dbname := "friendly_game"    

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, hostname, dbname)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	fmt.Println("Connected to MySQL database successfully!")
}