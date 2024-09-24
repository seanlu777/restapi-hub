// db/hub.go
package db

import (
	"gorm.io/gorm/clause"
)

// create Shipment
func CreateShipment(data *Shipment) error {
	if result := DB.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}

// retrieves all hubs from the database
func GetAllHubs() ([]HubHistory, error) {
	var histories []HubHistory

	result := DB.Raw(`
		SELECT * FROM hub_histories AS h
		WHERE h.id IN (
			SELECT MAX(id)
			FROM hub_histories
			GROUP BY gateway_id	
		)
	`).Scan(&histories)

	if result.Error != nil {
		return nil, result.Error
	}

	return histories, nil
}

// retrieves hub history by GatewayID
func GetHubHistories(gatewayID string) ([]HubHistory, error) {
	var history []HubHistory

	if result := DB.Where("gateway_id = ?", gatewayID).Find(&history); result.Error != nil {
		return nil, result.Error
	}
	return history, nil
}

// retrieves device list by its GatewayID
func GetDeviceList(gatewayID string) ([]DeviceList, error) {
	var a2tbDevices []A2TB
	var r2b2Devices []R2B2
	var r2t8Devices []R2T8

	var deviceList []DeviceList

	if err := DB.Raw(`
		SELECT * FROM a2_tbs AS a
		WHERE a.id IN (
			SELECT MAX(id)
			FROM a2_tbs
			WHERE gateway_id = ?
			GROUP BY device_id
		)
	`, gatewayID).Scan(&a2tbDevices).Error; err != nil {
		return nil, err
	}

	for _, device := range a2tbDevices {
		deviceList = append(deviceList, DeviceList{
			DeviceType: "A2TB",
			DeviceID:   device.DeviceID,
		})
	}

	if err := DB.Raw(`
		SELECT * FROM r2_b2 AS r
		WHERE r.id IN (
			SELECT MAX(id)
			FROM r2_b2
			WHERE gateway_id = ?
			GROUP BY device_id
		)
	`, gatewayID).Scan(&r2b2Devices).Error; err != nil {
		return nil, err
	}

	for _, device := range r2b2Devices {
		deviceList = append(deviceList, DeviceList{
			DeviceType: "R2B2",
			DeviceID:   device.DeviceID,
		})
	}

	if err := DB.Raw(`
		SELECT * FROM r2_t8 AS r
		WHERE r.id IN (
			SELECT MAX(id)
			FROM r2_t8
			WHERE gateway_id = ?
			GROUP BY device_id
		)	
	`, gatewayID).Scan(&r2t8Devices).Error; err != nil {
		return nil, err
	}

	for _, device := range r2t8Devices {
		deviceList = append(deviceList, DeviceList{
			DeviceType: "R2T8",
			DeviceID:   device.DeviceID,
		})
	}

	return deviceList, nil
}

// retrieves device data by its GatewayID
func GetDeviceData(gatewayID string, deviceType string, deviceID string) (map[string]interface{}, error) {
	deviceData := make(map[string]interface{})

	switch deviceType {
	case "A2TB":
		var a2tbDevices []A2TB
		if err := DB.Where("gateway_id = ? and device_id = ?", gatewayID, deviceID).
			Order("record_time DESC").
			Find(&a2tbDevices).Error; err != nil {
			return nil, err
		}
		deviceData["DeviceType"] = deviceType
		deviceData["Data"] = a2tbDevices

	case "R2B2":
		var r2b2Devices []R2B2
		if err := DB.Where("gateway_id = ? and device_id = ?", gatewayID, deviceID).
			Order("record_time DESC").
			Find(&r2b2Devices).Error; err != nil {
			return nil, err
		}
		deviceData["DeviceType"] = deviceType
		deviceData["Data"] = r2b2Devices

	case "R2T8":
		var r2t8Devices []R2T8
		if err := DB.Where("gateway_id = ? and device_id = ?", gatewayID, deviceID).
			Order("record_time DESC").
			Find(&r2t8Devices).Error; err != nil {
			return nil, err
		}
		deviceData["DeviceType"] = deviceType
		deviceData["Data"] = r2t8Devices
	}
	return deviceData, nil
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

// create A2TB
func CreateA2TB(data *A2TB) error {
	if result := DB.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}

// create R2B2
func CreateR2B2(data *R2B2) error {
	if result := DB.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}

// create R2T8
func CreateR2T8(data *R2T8) error {
	if result := DB.Create(data); result.Error != nil {
		return result.Error
	}
	return nil
}
