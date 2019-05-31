package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	PORT := "2000" // Default port 2000

	router := gin.Default()

	db := initDB() // Initialize the DB connection.
	
	router.POST("/login", verifyUser(db))			// Verify users login info and return a login token if successful.
	router.POST("/users/create", createUser(db))	// Create a user in the database.

	router.Run(":" + PORT)
}