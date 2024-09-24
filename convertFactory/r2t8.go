package convertFactory

import (
	"fmt"
	"gin-rest-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertR2T8(data string, recordTime int, gatewayID string) {

	DeviceID := data[0:6]

	Temp, err := strconv.ParseInt(data[6:8], 16, 64)
	if err != nil {
		return
	}
	TempF, err := strconv.ParseInt(data[8:10], 16, 64)
	if err != nil {
		return
	}
	Tempstr := strconv.FormatInt(Temp, 10) + "." + strconv.FormatInt(TempF, 10)
	Temperature, err := strconv.ParseFloat(Tempstr, 64)
	if err != nil {
		return
	}

	AsixX := "N/A"
	AsixY := "N/A"
	AsixZ := "N/A"

	BatteryLevel, err := strconv.ParseInt(data[34:36], 16, 64)
	if err != nil {
		return
	}
	// Debug
	// fmt.Println("DeviceID: ", DeviceID)
	// fmt.Println("Temperature: ", Temperature)
	// fmt.Println("AsixX: ", AsixX)
	// fmt.Println("AsixY: ", AsixY)
	// fmt.Println("AsixZ: ", AsixZ)
	// fmt.Println("BatteryLevel: ", BatteryLevel)

	r2t8Struct := db.R2T8{
		RecordTime:   recordTime,
		GatewayID:    gatewayID,
		DeviceID:     DeviceID,
		Temperature:  float32(Temperature),
		AsixX:        AsixX,
		AsixY:        AsixY,
		AsixZ:        AsixZ,
		BatteryLevel: int(BatteryLevel),
	}

	tx := db.DB.Begin()
	if err := db.CreateR2T8(&r2t8Struct); err != nil {
		fmt.Println(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

}
