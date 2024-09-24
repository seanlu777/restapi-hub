package convertFactory

import (
	"encoding/json"
	"fmt"
	"gin-rest-api/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConvertDeviceData(c *gin.Context, hub db.Hubs) {

	// Create gateway history
	history := db.HubHistory{
		GatewayID:     hub.GatewayID,
		LteRssi:       hub.LteRssi,
		WifiRssi:      hub.WifiRssi,
		SatelliteQty:  hub.SatelliteQty,
		Lng:           hub.Lng,
		Lat:           hub.Lat,
		Timestamp:     hub.Timestamp,
		MovementState: hub.MovementState,
		Light:         hub.Light,
		BatteryLevel:  hub.BatteryLevel,
	}

	sensorGJson, err := json.Marshal(hub.SensorG)
	if err != nil {
		log.Println("Error marshaling SensorG:", err)
		return
	}
	history.SensorG = sensorGJson
	tx := db.DB.Begin()

	if err := db.CreateHubHistory(&history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	records := hub.Records
	for i := range records {
		name := records[i].Name
		switch name {
		case "A2TB":
			fmt.Println("I'm A2TB; raw: ", records[i].RawData)
			ConvertA2TB(records[i].RawData, records[i].RecordTime, hub.GatewayID)
		case "R2B2":
			fmt.Println("I'm R2B2; raw: ", records[i].RawData)
			ConvertR2B2(records[i].RawData, records[i].RecordTime, hub.GatewayID)
		case "R2T8":
			fmt.Println("I'm R2T8; raw: ", records[i].RawData)
			ConvertR2T8(records[i].RawData, records[i].RecordTime, hub.GatewayID)
		}
	}
}
