package convertFactory

import (
	"fmt"
)

func ConvertR2B2(data string, recordTime int, gatewayID string) {

	DeviceID := data[0:6]
	Temperature := 0
	// Asix X
	xSign := func() string {
		if data[11] == 1 {
			return "-"
		}
		return ""
	}
	AsixX := xSign() + data[13:14] + "." + data[15:16] + data[17:18]

	// Asix Y
	ySign := func() string {
		if data[19] == 1 {
			return "-"
		}
		return ""
	}
	AsixY := ySign() + data[21:22] + "." + data[23:24] + data[25:26]

	// AsixZ
	zSign := func() string {
		if data[27] == 1 {
			return "-"
		}
		return ""
	}
	AsixZ := zSign() + data[29:30] + "-" + data[31:32] + data[33:34]

	// Battery level
	BatteryLevel := data[35:37]
	// Debug
	fmt.Println(DeviceID)
	fmt.Println(Temperature)
	fmt.Println(AsixX)
	fmt.Println(AsixY)
	fmt.Println(AsixZ)
	fmt.Println(BatteryLevel)
}
