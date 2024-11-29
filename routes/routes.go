package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// registercontroller "friendly-game/controllers"
	registercontroller "friendly-game/controllers"
)

func Registers() {
	fmt.Println("Registering routes...")
	router := gin.Default()
	router.GET("/register", registercontroller.Register)
	router.GET("/", func (c *gin.Context)  {
		c.String(200, "Welcome to Friendly Game!")	
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // No Content
	})	
	router.Run(":8080")
}