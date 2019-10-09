package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MySmartFarm/mysmartfarm_api/database"
	"github.com/MySmartFarm/mysmartfarm_api/functions"
	"github.com/MySmartFarm/mysmartfarm_api/models"
	"github.com/gin-gonic/gin"
)

func SoilMoistureSensorGet(c *gin.Context) {
	status, valid := functions.IsAuthorized(c.Request.Header, true)
	if status != "Success" && !valid {
		c.JSON(401, status)
		return
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}

	query := `
		SELECT * FROM soil_moisture_sensor WHERE sensor_id = '%d'
	`

	res, err := database.Query(fmt.Sprintf(query, ID))
	if err != nil {
		log.Println(err)
	}

	soilMoistureSensor := []models.SoilMoistureSensor{}

	fmt.Println(len(res[0].Series))
	if len(res[0].Series) < 1 {
		c.JSON(400, "ID not found data")
		return
	}

	for _, row := range res[0].Series[0].Values {
		soilMoist := models.SoilMoistureSensor{}
		soilMoist.Time, err = time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		soilMoist.MoistureValue, _ = strconv.Atoi(string(row[1].(json.Number)))
		soilMoist.SensorId, _ = strconv.Atoi(row[2].(string))
		soilMoist.StatusAlert, _ = strconv.Atoi(string(row[3].(json.Number)))

		soilMoistureSensor = append(soilMoistureSensor, soilMoist)
	}

	response := soilMoistureSensor
	c.JSON(200, response)

}

func SoilMoistureSensorGetAll(c *gin.Context) {
	status, valid := functions.IsAuthorized(c.Request.Header, true)

	if status != "Success" && !valid {
		c.JSON(401, status)
		return
	}

	query := `
		SELECT * FROM soil_moisture_sensor
	`

	res, err := database.Query(fmt.Sprintf(query))
	if err != nil {
		log.Println(err)
	}

	soilMoistureSensor := []models.SoilMoistureSensor{}

	fmt.Println(len(res[0].Series))
	if len(res[0].Series) < 1 {
		c.JSON(400, "ID not found data")
		return
	}

	for _, row := range res[0].Series[0].Values {
		soilMoist := models.SoilMoistureSensor{}
		soilMoist.Time, err = time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		soilMoist.MoistureValue, _ = strconv.Atoi(string(row[1].(json.Number)))
		soilMoist.SensorId, _ = strconv.Atoi(row[2].(string))
		soilMoist.StatusAlert, _ = strconv.Atoi(string(row[3].(json.Number)))

		soilMoistureSensor = append(soilMoistureSensor, soilMoist)
	}

	response := soilMoistureSensor
	c.JSON(200, response)

}

func SoilMoistureSensorCreate(c *gin.Context) {
	status, valid := functions.IsAuthorized(c.Request.Header, true)
	if status != "Success" && !valid {
		c.JSON(401, status)
		return
	}

	soilMoistureSensor := models.SoilMoistureSensor{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)

	err := json.Unmarshal(buf.Bytes(), &soilMoistureSensor)
	if err != nil {
		functions.NotifyToLine("SoilMoistureSensor : Format or data invalid")
		c.String(400, "")
		return
	}

	seriseName := "soil_moisture_sensor"
	soilMoistureSensor.Time = time.Now()
	soilMoistureSensor.MoistureValue = 1025 - soilMoistureSensor.MoistureValue
	if soilMoistureSensor.MoistureValue < 100 {
		soilMoistureSensor.StatusAlert = 1
	} else if soilMoistureSensor.MoistureValue < 200 {
		soilMoistureSensor.StatusAlert = 2
	} else if soilMoistureSensor.MoistureValue > 700 {
		soilMoistureSensor.StatusAlert = 4
	} else {
		soilMoistureSensor.StatusAlert = 3
	}

	// Create a point and add to batch
	tags := map[string]string{"sensor_id": fmt.Sprint(soilMoistureSensor.SensorId)}
	fields := map[string]interface{}{
		"moisture_value": soilMoistureSensor.MoistureValue,
		"status_alert":   soilMoistureSensor.StatusAlert,
	}

	database.Insert(seriseName, tags, fields, soilMoistureSensor.Time)
	response := soilMoistureSensor
	c.JSON(200, response)

}
