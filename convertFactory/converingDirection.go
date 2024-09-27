package convertFactory

import (
	"encoding/json"
	"fmt"
	"gin-rest-api/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConvertDeviceData(c *gin.Context, gateway db.Gateway) {

	// Create gateway history
	history := db.GatewayHistory{
		GatewayID:     gateway.GatewayID,
		LteRssi:       gateway.LteRssi,
		WifiRssi:      gateway.WifiRssi,
		SatelliteQty:  gateway.SatelliteQty,
		Lng:           gateway.Lng,
		Lat:           gateway.Lat,
		Timestamp:     gateway.Timestamp,
		MovementState: gateway.MovementState,
		Light:         gateway.Light,
		BatteryLevel:  gateway.BatteryLevel,
	}

	sensorGJson, err := json.Marshal(gateway.SensorG)
	if err != nil {
		log.Println("Error marshaling SensorG:", err)
		return
	}
	history.SensorG = sensorGJson
	tx := db.DB.Begin()
	// Attempt to create the shipment along with its associated devices
	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit the transaction if everything is okay
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	records := gateway.Records
	for i := range records {
		name := records[i].Name
		switch name {
		case "A2TB":
			fmt.Println("I'm A2TB; raw: ", records[i].RawData)
			ConvertA2TB(records[i].RawData, records[i].RecordTime, gateway.GatewayID)
		case "R2B2":
			fmt.Println("I'm R2B2; raw: ", records[i].RawData)
			ConvertR2B2(records[i].RawData, records[i].RecordTime, gateway.GatewayID)
		case "R2T8":
			fmt.Println("I'm R2T8; raw: ", records[i].RawData)
			ConvertR2T8(records[i].RawData, records[i].RecordTime, gateway.GatewayID)
		}
	}
}
