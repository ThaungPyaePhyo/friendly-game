package migrate

import (
	"fmt"
	"database/sql"
	_"friendly-game/database"
)

func Migrating(db *sql.DB) {
	fmt.Println("Migrating the database...")	
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			unique_id VARCHAR(255) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS friendships (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			friend_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (friend_id) REFERENCES users(id)
		);`,
		`CREATE TABLE IF NOT EXISTS games (
			id INT AUTO_INCREMENT PRIMARY KEY,
			player1_id INT NOT NULL,
			player2_id INT NOT NULL,
			game_status VARCHAR(50),
			FOREIGN KEY (player1_id) REFERENCES users(id),
			FOREIGN KEY (player2_id) REFERENCES users(id)
		);`,
	}

	for _, query := range migrations {
		_, err := db.Exec(query)
		if  err != nil {
			fmt.Printf("Error running migration: %v\n", err)
		} else{
			fmt.Println("Migration ran successfully!")	
		}
	}
}