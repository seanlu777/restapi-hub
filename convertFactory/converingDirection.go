package convertFactory

import (
	"fmt"
	"gin-rest-api/db"
)

func ConvertDeviceData(hub db.Hubs) {

	records := hub.Records
	for i := range records {
		name := records[i].Name
		switch name {
		case "A2TB":
			fmt.Println("I'm A2TB; raw: ", records[i].RawData)
			ConvertA2TB(records[i].RawData, hub.GatewayID)
		case "R2B2":
			fmt.Println("I'm R2B2; raw: ", records[i].RawData)
		case "R2T8":
			fmt.Println("I'm R2T8; raw: ", records[i].RawData)
		}
	}
}
