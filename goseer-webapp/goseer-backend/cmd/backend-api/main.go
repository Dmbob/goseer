package main

import (
	"os"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string `json:"serverport"`
	DBUser string `"json:"dbuser"`
	DBPass string `json:"dbpass"`
	DBAddress string `json:"dbaddress"`
	DBPort string `json:"dbport"`
	DBName string `json"dbname"`
}

func main() {
	var config Config

	file, err := os.Open("./config/config.json")

	if err != nil { panic(err) }

	decoder := json.NewDecoder(file)
	decoder.Decode(&config)

	PORT := config.Port

	router := gin.Default()

	db := initDB(&config) // Initialize the DB connection.
	
	router.POST("/login", verifyUser(db))			// Verify users login info and return a login token if successful.
	router.POST("/users/create", createUser(db))	// Create a user in the database.

	router.Run(":" + PORT)
}