// db/models.go
package db

import (
	"gorm.io/gorm"
)

// Define the Hubs struct with GORM tags
type Hubs struct {
	GatewayID      string    `gorm:"unique; size:20; not null; index" json:"GatewayID"`
	LteRssi        int       `json:"LteRssi"`
	WifiRssi       int       `json:"WifiRssi"`
	SatelliteQty   int       `json:"SatelliteQty"`
	Lng            float32   `json:"Lng"`
	Lat            float32   `json:"Lat"`
	Timestamp      int       `json:"Timestamp"`
	SensorG        []float32 `json:"SensorG"`
	MovementState  int       `json:"MovementState"`
	Light          int       `json:"Light"`
	BatteryLevel   int       `json:"BatteryLevel"`
	ChargingStatus int       `json:"ChargingStatus"`
	Records        []Records `json:"Records"`
}

// Define the Records struct with GORM tags
type Records struct {
	GatewayID  string `json:"GatewayID"`
	Name       string `gorm:"size: 20; not null" json:"Name"`
	RecordTime int    `gorm:"RecordTime; not null" json:"RecordTime"`
	RawData    string `json:"RawData"`
}

// Define the hub history for tracking paths
type HubHistory struct {
	gorm.Model
	GatewayID      string  `gorm:"size: 20; not null; index" json:"GatewayID"`
	LteRssi        int     `json:"LteRssi"`
	WifiRssi       int     `json:"WifiRssi"`
	SatelliteQty   int     `json:"SatelliteQty"`
	Lng            float32 `json:"Lng"`
	Lat            float32 `json:"Lat"`
	Timestamp      int     `json:"Timestamp"`
	SensorG        string  `json:"SensorG"`
	MovementState  int     `json:"MovementState"`
	Light          int     `json:"Light"`
	BatteryLevel   int     `json:"BatteryLevel"`
	ChargingStatus int     `json:"ChargingStatus"`
}

// Define A2TB
type A2TB struct {
	gorm.Model
	RecordTime   int     `gorm:"RecordTime; not null" json:"RecordTime"`
	GatewayID    string  `gorm:"size: 20" json:"GatewayID"`
	DeviceID     string  `json:"DeviceID"`
	Temperature  float32 `json:"Temperature"`
	Pressure     float32 `json:"Pressure"`
	TagStatus    string  `json:"TagStatus"`
	BatteryLevel int     `json:"BatteryLevel"`
	TImestamp    int     `json:"Timestemp"`
	TXPower      int     `json:"TXPower"`
	ReserveData  string  `json:"ReserveData"`
}

// Define R2B2
type R2B2 struct {
	gorm.Model
	RecordTime   int     `gorm:"RecordTime; not null" json:"RecordTime"`
	GatewayID    string  `gorm:"size: 20" json:"GatewayID"`
	DeviceID     string  `json:"DeviceID"`
	Temperature  float32 `json:"Temperature"`
	AsixX        float32 `json:"AsixX"`
	AsixY        float32 `json:"AsixY"`
	AsixZ        float32 `json:"AsixZ"`
	BatteryLevel int     `json:"BatteryLevel"`
}

// Define T8
type R2T8 struct {
	gorm.Model
	RecordTime   int     `gorm:"RecordTime; not null" json:"RecordTime"`
	GatewayID    string  `gorm:"size: 20" json:"GatewayID"`
	DeviceID     string  `json:"DeviceID"`
	Temperature  float32 `json:"Temperature"`
	AsixX        float32 `json:"AsixX"`
	AsixY        float32 `json:"AsixY"`
	AsixZ        float32 `json:"AsixZ"`
	BatteryLevel int     `json:"BatteryLevel"`
}
