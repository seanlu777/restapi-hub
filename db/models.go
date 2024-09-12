// db/models.go
package db

import "gorm.io/gorm"

// Define the Hubs struct with GORM tags
type Hubs struct {
	gorm.Model
	GatewayID    string    `gorm:"unique; size:20; not null; index" json:"GatewayID"`
	LteRssi      int       `json:"LteRssi"`
	WifiRssi     int       `json:"WifiRssi"`
	SatelliteQty int       `json:"SatelliteQty"`
	Lng          float32   `json:"Lng"`
	Lat          float32   `json:"Lat"`
	Timestamp    int       `json:"Timestamp"`
	Records      []Records `gorm:"foreignKey:GatewayID;references:GatewayID;constraint:OnDelete:CASCADE" json:"Records"`
}

// Define the Records struct with GORM tags
type Records struct {
	gorm.Model
	GatewayID  string `gorm:"size: 20; not null; index" json:"GatewayID"` // fk: GatewayID string
	DeviceID   string `gorm:"size: 20; not null" json:"DeviceID"`
	Name       string `gorm:"size: 20; not null" json:"Name"`
	RecordTime string `gorm:"RecordTime; not null" json:"RecordTime"`
	RawData    string `json:"RawData"`
}

// Define the hub history for tracking paths
type HubHistory struct {
	gorm.Model
	GatewayID    string  `gorm:"size: 20; not null; index" json:"GatewayID"`
	LteRssi      int     `json:"LteRssi"`
	WifiRssi     int     `json:"WifiRssi"`
	SatelliteQty int     `json:"SatelliteQty"`
	Lng          float32 `json:"Lng"`
	Lat          float32 `json:"Lat"`
	Timestamp    int     `json:"Timestamp"`
}
