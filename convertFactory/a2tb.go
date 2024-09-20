package convertFactory

import (
	"fmt"
	"strconv"
)

func ConvertA2TB(data string, gatewayID string) {

	RepeaterID := data[0:6]
	DeviceID := data[6:12]

	// Temperature
	Temp, err := strconv.ParseInt(data[12:14], 16, 64)
	if err != nil {
		return
	}
	TempF, err := strconv.ParseInt(data[14:16], 16, 64)
	if err != nil {
		return
	}
	Temperature := strconv.FormatInt(Temp, 10) + "." + strconv.FormatInt(TempF, 10)

	// Pressure
	Pressure := data[16:22]

	// Status
	TagStatus := data[22:24]
	switch TagStatus {
	case "6c":
		TagStatus = "Connect"
	case "6d":
		TagStatus = "Open"
	case "6e":
		TagStatus = "Temperature Alert"
	case "6a":
		TagStatus = "Pressure Alert"
	case "aa":
		TagStatus = "Button Alert"
	case "ba":
		TagStatus = "Battery Low Alert"
	}

	// Battery Level
	BatteryLevel, err := strconv.ParseInt(data[24:26], 16, 64)
	if err != nil {
		return
	}

	// Timestemp
	Timestamp, err := strconv.ParseInt(data[26:32], 16, 64)
	if err != nil {
		return
	}

	// TX Power
	TXPower, err := strconv.ParseInt(data[32:34], 16, 64)
	if err != nil {
		return
	}

	// ReserveData
	ReserveData, err := strconv.ParseInt(data[34:36], 16, 64)
	if err != nil {
		return
	}

	fmt.Println(RepeaterID)
	fmt.Println(DeviceID)
	fmt.Println(Temperature)
	fmt.Println(Pressure)
	fmt.Println(TagStatus)
	fmt.Println(BatteryLevel)
	fmt.Println(Timestamp)
	fmt.Println(TXPower)
	fmt.Println(ReserveData)
	// if err := db.CreateA2TB(); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
}
