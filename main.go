package main

import (
	"fmt"

	"gin-rest-api/db"

	"gin-rest-api/convertFactory"

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
	router.GET("/api/getGateways", getGateways)
	router.GET("/api/getGatewayHistory", getGatewayHistory)
	router.GET("/api/getDeviceList", getDeviceList)
	router.GET("/api/getDeviceData", getDeviceData)
	router.POST("/api/pushRecord", pushRecord)
	router.POST("/api/createShipment", createShipment)
	router.PUT("/api/updateShipment", updateShipment)

	// Start the server on the port 8080
	router.Run(":8080")
}

// handler function to check the API's health status
func apiHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// handler function to get all hubs from the database and return them as JSON response
func getGateways(c *gin.Context) {
	gateways, err := db.GetAllGateways()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "OK", "Gateways": gateways})
}

// handler function to get a single hub by its GatewayID
func getGatewayHistory(c *gin.Context) {
	id := c.Query("gatewayID") // Directly use the ID as a String
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": "gatewayID is required"})
	}
	gateway, err := db.GetGatewayHistories(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid gatewayID"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "OK", "Gateway": gateway})
}

// handler function to get Device List
func getDeviceList(c *gin.Context) {
	id := c.Query("gatewayID")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": "gatewayID is required"})
	}
	deviceList, err := db.GetDeviceList(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error: ": "Invalid gatewayID"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"Status": "OK", "Device": deviceList})
}

// handler function to get Device data
func getDeviceData(c *gin.Context) {
	gatewayID := c.Query("gatewayID")
	if gatewayID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": "gatewayID is required"})
	}

	deviceType := c.Query("deviceType")
	if deviceType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": "deviceType is required"})
	}

	deviceID := c.Query("deviceID")
	if deviceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": "deviceID is required"})
	}

	device, err := db.GetDeviceData(gatewayID, deviceType, deviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"Status": "OK", "Device": device})
}

// Handler function to create a new hub and record in the database
func pushRecord(c *gin.Context) {
	var data db.Gateway

	if err := c.BindJSON(&data); err != nil {
		fmt.Printf("BindJSON Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	convertFactory.ConvertDeviceData(c, data)
	fmt.Printf("data: %+v\n", data)
	fmt.Println()

	c.JSON(http.StatusCreated, data)
}

// Create shipment and device list
func createShipment(c *gin.Context) {
	var shipment db.Shipment

	if err := c.BindJSON(&shipment); err != nil {
		fmt.Printf("BindJSON Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	convertFactory.SaveShipment(c, shipment)

	c.JSON(http.StatusCreated, gin.H{"Status": "Created", "Shipment": shipment})
}

// Update shipment
func updateShipment(c *gin.Context) {
	var shipment db.Shipment

	if err := c.BindJSON(&shipment); err != nil {
		fmt.Printf("BindJSON Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	convertFactory.UpdateShipment(c, shipment)

	c.JSON(http.StatusOK, gin.H{"Status": "Update succeeded", "Shipment": shipment})
}
