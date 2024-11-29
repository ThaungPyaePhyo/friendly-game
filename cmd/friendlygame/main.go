package main

import (
	"fmt"
	// "friendly-game/database"
	// migrate "friendly-game/migrations"
	"friendly-game/routes"
)

func main() {
	fmt.Println("Hello World")
	// database.Connect()
	// defer database.DB.Close()
	// migrate.Migrating(database.DB)
	routes.Registers()

}