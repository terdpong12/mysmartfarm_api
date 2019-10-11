package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MySmartFarm/mysmartfarm_api/constants"
	"github.com/MySmartFarm/mysmartfarm_api/database"
	"github.com/MySmartFarm/mysmartfarm_api/functions"
	"github.com/MySmartFarm/mysmartfarm_api/models"
	"github.com/gin-gonic/gin"
)

func GetSoilMoistureSensor(c *gin.Context) {
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
		SELECT * FROM ` + constants.SeriesNameSoilMoistureSensor + ` WHERE sensor_id = '%d' ORDER BY time DESC LIMIT 180
	`
	res, err := database.Query(fmt.Sprintf(query, ID))
	if err != nil {
		log.Println(err)
	}

	soilMoistureSensor := []models.SoilMoistureSensor{}
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
		soilMoist.SensorId, _ = strconv.Atoi(row[1].(string))
		soilMoist.StatusAlert, _ = strconv.Atoi(string(row[2].(json.Number)))
		soilMoist.Value, _ = strconv.Atoi(string(row[3].(json.Number)))
		soilMoistureSensor = append(soilMoistureSensor, soilMoist)
	}

	response := soilMoistureSensor
	c.JSON(200, response)

}

func GetListSoilMoistureSensor(c *gin.Context) {
	status, valid := functions.IsAuthorized(c.Request.Header, true)

	if status != "Success" && !valid {
		c.JSON(401, status)
		return
	}

	query := `
		SELECT * FROM ` + constants.SeriesNameSoilMoistureSensor + ` ORDER BY time DESC LIMIT 180
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
		soilMoist.SensorId, _ = strconv.Atoi(row[1].(string))
		soilMoist.StatusAlert, _ = strconv.Atoi(string(row[2].(json.Number)))
		soilMoist.Value, _ = strconv.Atoi(string(row[3].(json.Number)))

		soilMoistureSensor = append(soilMoistureSensor, soilMoist)
	}

	response := soilMoistureSensor
	c.JSON(200, response)

}

func CreateSoilMoistureSensor(c *gin.Context) {
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

	soilMoistureSensor.Time = time.Now()
	soilMoistureSensor.Value = 1023 - soilMoistureSensor.Value
	if soilMoistureSensor.Value < 100 {
		soilMoistureSensor.StatusAlert = 1
	} else if soilMoistureSensor.Value < 200 {
		soilMoistureSensor.StatusAlert = 2
	} else if soilMoistureSensor.Value > 700 {
		soilMoistureSensor.StatusAlert = 4
	} else {
		soilMoistureSensor.StatusAlert = 3
	}

	// Create a point and add to batch
	tags := map[string]string{"sensor_id": fmt.Sprint(soilMoistureSensor.SensorId)}
	fields := map[string]interface{}{
		"value":        soilMoistureSensor.Value,
		"status_alert": soilMoistureSensor.StatusAlert,
	}

	database.Insert(constants.SeriesNameSoilMoistureSensor, tags, fields, soilMoistureSensor.Time)
	response := soilMoistureSensor
	c.JSON(200, response)

}
