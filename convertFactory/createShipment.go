package convertFactory

import (
	"gin-rest-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveShipment(c *gin.Context, shipment db.Shipment) {

	tx := db.DB.Begin()
	// Attempt to create the shipment
	if err := tx.Create(&shipment).Error; err != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit the transaction if everything is okay
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
