package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	router := gin.Default()

	// Define API Routes
	v1 := router.Group("/api/v1", gin.BasicAuth(gin.Accounts{
		"bob": "test",
	}))
	{
		v1.GET("/temperatures", fetchSystemTemps)	// Grab System Temperatures
	}

	router.Run()
}
