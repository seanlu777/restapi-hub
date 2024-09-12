package main

import (
	"fmt"

	"gin-rest-api/db"

	"gorm.io/gorm"

	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=localhost user=test1 password=test123 dbname=tydenbrooks_hub"
	err := db.Initialize(dsn)
	if err != nil {
		log.Fatal(err)
	}
	// Set Gin to release mode to minimize logging
	gin.SetMode(gin.ReleaseMode)

	// Intitialize Gin router
	router := gin.Default()

	// Configure trusted proxies

	if err := router.SetTrustedProxies([]string{"0.0.0.0/0"}); err != nil {
		panic("Failed to set trusted proxies: " + err.Error())
	}

	// Define routes
	router.GET("/api/health", apiHealth)
	router.GET("/api/hubs", getHubs)
	router.GET("/api/hub/:gatewayID", getHub)
	router.POST("/api/pushRecord", pushRecord)

	// Start the server on the port 8080
	router.Run(":8080")
}

// handler function to check the API's health status
func apiHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// handler function to get all hubs from the database and return them as JSON response
func getHubs(c *gin.Context) {
	hubs, err := db.GetAllHubs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hubs)
}

// handler function to get a single hub by its GatewayID
func getHub(c *gin.Context) {
	id := c.Param(":gatewayID") // Directly use the ID as a String
	hub, err := db.GetHub(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid gateway ID"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, hub)
}

// Handler function to create a new hub and record in the database
func pushRecord(c *gin.Context) {

	var data db.Hubs

	if err := c.BindJSON(&data); err != nil {
		fmt.Printf("BindJSON Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := range data.Records {
		data.Records[i].GatewayID = data.GatewayID
	}
	fmt.Printf("data: %+v\n", data)
	fmt.Println()

	history := db.HubHistory{
		GatewayID:    data.GatewayID,
		LteRssi:      data.LteRssi,
		WifiRssi:     data.WifiRssi,
		SatelliteQty: data.SatelliteQty,
		Lng:          data.Lng,
		Lat:          data.Lat,
		Timestamp:    data.Timestamp,
	}

	if err := db.CreateHubHistory(&history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx := db.DB.Begin()
	if err := db.CreateData(&data); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	c.JSON(http.StatusCreated, data)
}
