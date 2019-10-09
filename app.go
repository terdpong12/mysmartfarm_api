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
	r.GET("/sensor/soil_moisture", controllers.SoilMoistureSensorGetAll)
	r.GET("/sensor/soil_moisture/:id", controllers.SoilMoistureSensorGet)
	r.POST("/sensor/soil_moisture", controllers.SoilMoistureSensorCreate)

	r.Run() //listen and serve on 0.0.0.0:8080
}
