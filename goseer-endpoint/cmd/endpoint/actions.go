package main

import (
	"fmt"
	"strings"
	"github.com/ssimunic/gosensors"

)

func fetchSystemTemps() []map[string]string {
	sensors, err := gosensors.NewFromSystem()
	var data []map[string]string
	var tempString string

	if err != nil {
		fmt.Println(err)
	}

	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
			if key == "CPU" || key == "GPU" || strings.Contains(key, "Core") {
				returnValue := make(map[string]string)

				startIndex := strings.Index(value, "+")
				endIndex := strings.Index(value, " ") - 3
				highStartIndex := strings.Index(value, "high") + len("high") + 4
				highEndIndex := strings.Index(value, ",") - 3
				criticalStartIndex := strings.Index(value, "crit") + len("crit") + 4
				criticalEndIndex := strings.Index(value, ")") - 3

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
				returnValue["highTemp"] = highValue
				returnValue["criticalTemp"] = criticalValue

				data = append(data, returnValue)
				tempString += key + ": " + value + "|"
			}
		}
	}

	return data

}
