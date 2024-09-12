// db/hub.go
package db

import (
	"gorm.io/gorm/clause"
)

// retrieves all hubs from the database
func GetAllHubs() ([]Hubs, error) {
	var hubs []Hubs

	if result := DB.Preload("Records").Find(&hubs); result.Error != nil {
		return nil, result.Error
	}
	return hubs, nil
}

// retrieves a hub by its GatewayID
func GetHub(gatewayID string) (*Hubs, error) {
	var hub Hubs

	if result := DB.Preload("Records").Where("gatewayID = ?", gatewayID).First(&hub); result.Error != nil {
		return nil, result.Error
	}
	return &hub, nil
}

// creates a new hub in the database
func CreateData(data *Hubs) error {

	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "gateway_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"lte_rssi", "wifi_rssi", "satellite_qty", "lng", "lat", "timestamp"}),
	}).Create(data)

	return result.Error
}

// create hub history
func CreateHubHistory(data *HubHistory) error {
	if result := DB.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}
