// db/models.go
package db

import (
	"gorm.io/gorm"
)

// Define the shipments table
type Shipment struct {
	gorm.Model
	ShipmentID  string      `gorm:"unique; size:20; not null; index" json:"ShipmentID"`
	CreatorName string      `json:"CreatorName"`
	Destination Destination `gorm:"type:json" json:"Destination"` // Country, City, Area, Address
	GatewayType string      `json:"GatewayType"`
	GatewayID   string      `json:"GatewayID"`
	Status      string      `json:"Status"`
	Device      []Device    `gorm:"foreignKey:ShipmentID; references:ShipmentID" json:"Device"`
}

// Define the destination struct
type Destination struct {
	Country string `json:"Country"`
	City    string `json:"City"`
	Area    string `json:"Area"`
	Address string `json:"Address"`
}

// Define the device struct
type Device struct {
	gorm.Model
	ShipmentID string `gorm:"not null;" json:"ShipmentID"`
	DeviceType string `json:"DeviceType"`
	DeviceID   string `gorm:"not null; unique" json:"DeviceID"` // Unique DeviceID
}

// Define the hubs
type Gateway struct {
	GatewayID      string    `json:"GatewayID"`
	LteRssi        int       `json:"LteRssi"`
	WifiRssi       int       `json:"WifiRssi"`
	SatelliteQty   int       `json:"SatelliteQty"`
	Lng            float32   `json:"Lng"`
	Lat            float32   `json:"Lat"`
	Timestamp      int       `json:"Timestamp"`
	SensorG        []float32 `gorm:"type:json" json:"SensorG"`
	MovementState  int       `json:"MovementState"`
	Light          int       `json:"Light"`
	BatteryLevel   int       `json:"BatteryLevel"`
	ChargingStatus int       `json:"ChargingStatus"`
	Records        []Records `json:"Records"`
}

// Define the records
type Records struct {
	Name       string `gorm:"size: 20; not null" json:"Name"`
	RecordTime int    `gorm:"RecordTime; not null" json:"RecordTime"`
	RawData    string `json:"RawData"`
}

// Define the hub history for tracking paths
type GatewayHistory struct {
	gorm.Model
	GatewayID      string  `gorm:"size: 20; not null; index" json:"GatewayID"`
	LteRssi        int     `json:"LteRssi"`
	WifiRssi       int     `json:"WifiRssi"`
	SatelliteQty   int     `json:"SatelliteQty"`
	Lng            float32 `json:"Lng"`
	Lat            float32 `json:"Lat"`
	Timestamp      int     `json:"Timestamp"`
	SensorG        []byte  `gorm:"type: json" json:"SensorG"`
	MovementState  int     `json:"MovementState"`
	Light          int     `json:"Light"`
	BatteryLevel   int     `json:"BatteryLevel"`
	ChargingStatus int     `json:"ChargingStatus"`
}

// Define A2TB
type A2TB struct {
	gorm.Model
	RecordTime   int     `gorm:"RecordTime; not null" json:"RecordTime"`
	RepeaterID   string  `json:"RepeaterID"`
	GatewayID    string  `gorm:"size: 20" json:"GatewayID"`
	DeviceID     string  `json:"DeviceID"`
	Temperature  float32 `json:"Temperature"`
	Pressure     float32 `json:"Pressure"`
	TagStatus    string  `json:"TagStatus"`
	BatteryLevel int     `json:"BatteryLevel"`
	Timestamp    int     `json:"Timestemp"`
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
	AsixX        string  `json:"AsixX"`
	AsixY        string  `json:"AsixY"`
	AsixZ        string  `json:"AsixZ"`
	BatteryLevel int     `json:"BatteryLevel"`
}

// Define T8
type R2T8 struct {
	gorm.Model
	RecordTime   int     `gorm:"RecordTime; not null" json:"RecordTime"`
	GatewayID    string  `gorm:"size: 20" json:"GatewayID"`
	DeviceID     string  `json:"DeviceID"`
	Temperature  float32 `json:"Temperature"`
	AsixX        string  `json:"AsixX"`
	AsixY        string  `json:"AsixY"`
	AsixZ        string  `json:"AsixZ"`
	BatteryLevel int     `json:"BatteryLevel"`
}

// Define DeviceList
type DeviceList struct {
	DeviceType string `json:"DeviceType"`
	DeviceID   string `json:"DeviceID"`
}

// Define DeviceData
type DeviceData struct {
	DeviceType string `json:"DeviceType"`
	A2TB       []A2TB `json:"A2TB"`
	R2B2       []R2B2 `json:"R2B2"`
	R2T8       []R2T8 `json:"R2T8"`
}
