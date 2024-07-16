package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"api/internal"
)

func main() {
	router := gin.Default()

	db := internal.ConnectDB()
	if db == nil {
		log.Println("failed to connect to database: ", db)
		return
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	router.GET("/send/:message", internal.HandleMessages)

	router.Run(":8080")
}
