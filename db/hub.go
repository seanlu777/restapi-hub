// db/hub.go
package db

// GetHubs retrieves all hubs from the database
func GetAllHubs() ([]Hubs, error) {
	var hubs []Hubs

	if result := DB.Preload("Records").Find(&hubs); result.Error != nil {
		return nil, result.Error
	}
	return hubs, nil
}

// GetHub retrieves a hub by its GatewayID
func GetHub(gatewayID string) (*Hubs, error) {
	var hub Hubs

	if result := DB.Preload("Records").Where("gatewayID = ?", gatewayID).First(&hub); result.Error != nil {
		return nil, result.Error
	}
	return &hub, nil
}

// CreateHub creates a new hub in the database
func CreateHub(hub *Hubs) error {
	if result := DB.Create(hub); result.Error != nil {
		return result.Error
	}
	return nil
}
func CreateRecord(record *Records) error {
	if result := DB.Create(record); result.Error != nil {
		return result.Error
	}
	return nil
}
