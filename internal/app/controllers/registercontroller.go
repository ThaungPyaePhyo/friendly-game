package registercontroller

import (
	// "fmt"
	// "net/http"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.String(200, "Welcome to Friendly Game!")	
	// c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
