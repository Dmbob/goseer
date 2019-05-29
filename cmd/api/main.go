package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	initDB()

	router := gin.Default()
	router.Use(cors.Default())
	// Define API Routes
	//v1 := router.Group("/api/v1", gin.BasicAuth(gin.Accounts{
	//	"bob": "test",
	//}))
	v1 := router.Group("/api/v1")
	{
		v1.GET("/temperatures", fetchSystemTemps)	// Grab System Temperatures
	}

	router.Run()
}
