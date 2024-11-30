package main

import (
	"fmt"
	"net/http"
	"log"
	// "friendly-game/routes"
	database "friendly-game/configs"
)

func main() {
	database.Connect()

	fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }

}