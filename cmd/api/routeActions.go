package main

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/ssimunic/gosensors"

)

func fetchSystemTemps(c *gin.Context) {
	sensors, err := gosensors.NewFromSystem()
	var data []string
	var tempString string

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		panic(err)
	}

	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
			if key == "CPU" || key == "GPU" || strings.Contains(key, "Core") {
				data = append(data, key, value)
				tempString += key + ": " + value + "|"
			}
		}
	}
	tempString = strings.TrimRight(tempString, "|")

	dbTemp := tempModel{Temps: tempString}

	db.Save(&dbTemp)

	c.JSON(http.StatusOK, gin.H{"temps": data})

}
