package main

import (
	"fmt"

	"github.com/MySmartFarm/mysmartfarm_api/controllers"
	"github.com/MySmartFarm/mysmartfarm_api/database"
	"github.com/MySmartFarm/mysmartfarm_api/functions"
	"github.com/gin-gonic/gin"
)

func main() {

	//http API
	fmt.Println(functions.GenerateJWT())
	database.CreateDatabase()
	r := gin.Default()
	r.GET("/getHello/", controllers.GetHello)
	r.GET("/getENV/", controllers.GetENV)

	//Soil Moisture Sensor
	r.GET("/sensor/soil_moisture", controllers.GetListSoilMoistureSensor)
	r.GET("/sensor/soil_moisture/:id", controllers.GetSoilMoistureSensor)
	r.POST("/sensor/soil_moisture", controllers.CreateSoilMoistureSensor)

	//LDR Photoresistor Sensor
	r.GET("/sensor/ldr_photoresistor", controllers.GetListLDRPhotoresistorSensor)
	r.GET("/sensor/ldr_photoresistor/:id", controllers.GetLDRPhotoresistorSensor)
	r.POST("/sensor/ldr_photoresistor", controllers.CreateLDRPhotoresistorSensor)

	r.Run() //listen and serve on 0.0.0.0:8080
}
