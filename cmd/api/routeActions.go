package main

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/ssimunic/gosensors"

)

func fetchSystemTemps(c *gin.Context) {
	sensors, err := gosensors.NewFromSystem()
	var data []map[string]string
	var tempString string

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		panic(err)
	}

	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
			if key == "CPU" || key == "GPU" || strings.Contains(key, "Core") {
				returnValue := make(map[string]string)

				startIndex := strings.Index(value, "+")
				endIndex := strings.Index(value, " ")
				highStartIndex := strings.Index(value, "high") + len("high") + 3
				highEndIndex := strings.Index(value, ",")
				criticalStartIndex := strings.Index(value, "crit") + len("crit") + 3
				criticalEndIndex := strings.Index(value, ")")

				newValue := ""
				highValue := ""
				criticalValue := ""

				if len(value) > 0 {
					newValue = value[startIndex:endIndex]
					highValue = value[highStartIndex:highEndIndex]
					criticalValue = value[criticalStartIndex:criticalEndIndex]
				}
				returnValue["adapter"] = key
				returnValue["currentTemp"] = newValue
				returnValue["maxTemp"] = highValue
				returnValue["criticalTemp"] = criticalValue

				data = append(data, returnValue)
				tempString += key + ": " + value + "|"
			}
		}
	}
	tempString = strings.TrimRight(tempString, "|")

	//dbTemp := tempModel{Temps: tempString}

	//db.Save(&dbTemp)

	c.JSON(http.StatusOK, gin.H{"temps": data})

}
