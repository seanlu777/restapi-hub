package convertFactory

import (
	"fmt"
	"gin-rest-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertA2TB(data string, RecordTime int, GatewayID string) {

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
	TempStr := strconv.FormatInt(Temp, 10) + "." + strconv.FormatInt(TempF, 10)
	Temperature, err := strconv.ParseFloat(TempStr, 32)
	if err != nil {
		return
	}
	// Pressure
	PressStr := data[16:22]
	Pressure, err := strconv.ParseFloat(PressStr, 32)
	if err != nil {
		return
	}
	// Status
	TagStatus := data[22:24]
	switch TagStatus {
	case "6c":
		TagStatus = "Connect"
	case "6d":
		TagStatus = "Open / Tampering Cut"
	case "6e":
		TagStatus = "Temperature Alert"
	case "6a":
		TagStatus = "Pressure Alert"
	case "ba":
		TagStatus = "Battery Low Alert"
	}
	// Battery Level
	BatteryLevel, err := strconv.ParseInt(data[24:26], 16, 64)
	if err != nil {
		return
	}
	// Timestemp
	Count, err := strconv.ParseInt(data[26:32], 16, 64)
	if err != nil {
		return
	}
	Timestamp := Count * 15
	// TX Power
	TXPower, err := strconv.ParseInt(data[32:34], 16, 64)
	if err != nil {
		return
	}
	// ReserveData
	ReserveData := data[34:36]

	// Debug
	// fmt.Println("RepeaterID: ", RepeaterID)
	// fmt.Println("DeviceID: ", DeviceID)
	// fmt.Println("Temperature: ", Temperature)
	// fmt.Println("Pressure: ", Pressure)
	// fmt.Println("TagStatus: ", TagStatus)
	// fmt.Println("BatteryLevel: ", BatteryLevel)
	// fmt.Println("Timestamp: ", Timestamp)
	// fmt.Println("TXPower: ", TXPower)
	// fmt.Println("ReserveData: ", ReserveData)

	// Saving to SQL
	a2tbStruct := db.A2TB{
		RecordTime:   RecordTime,
		RepeaterID:   RepeaterID,
		GatewayID:    GatewayID,
		DeviceID:     DeviceID,
		Temperature:  float32(Temperature),
		Pressure:     float32(Pressure),
		TagStatus:    TagStatus,
		BatteryLevel: int(BatteryLevel),
		Timestamp:    int(Timestamp),
		TXPower:      int(TXPower),
		ReserveData:  ReserveData,
	}

	tx := db.DB.Begin()
	if err := db.CreateA2TB(&a2tbStruct); err != nil {
		fmt.Println(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()
}
