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

func GetLDRPhotoresistorSensor(c *gin.Context) {
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
		SELECT * FROM ` + constants.SeriesNameLDRPhotoresistor + ` WHERE sensor_id = '%d'
	`
	res, err := database.Query(fmt.Sprintf(query, ID))
	if err != nil {
		log.Println(err)
	}

	LDRPhotoresistorSensor := []models.LDRPhotoresistorSensor{}
	if len(res[0].Series) < 1 {
		c.JSON(400, "ID not found data")
		return
	}

	for _, row := range res[0].Series[0].Values {
		LDRPhoto := models.LDRPhotoresistorSensor{}
		LDRPhoto.Time, err = time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		LDRPhoto.Value, _ = strconv.Atoi(string(row[1].(json.Number)))
		LDRPhoto.SensorId, _ = strconv.Atoi(row[2].(string))
		LDRPhoto.StatusAlert, _ = strconv.Atoi(string(row[3].(json.Number)))
		LDRPhotoresistorSensor = append(LDRPhotoresistorSensor, LDRPhoto)
	}

	response := LDRPhotoresistorSensor
	c.JSON(200, response)

}

func GetListLDRPhotoresistorSensor(c *gin.Context) {
	status, valid := functions.IsAuthorized(c.Request.Header, true)

	if status != "Success" && !valid {
		c.JSON(401, status)
		return
	}

	query := `
		SELECT * FROM ` + constants.SeriesNameLDRPhotoresistor + `  ORDER BY time DESC LIMIT 180
	`

	res, err := database.Query(fmt.Sprintf(query))
	if err != nil {
		log.Println(err)
	}

	LDRPhotoresistorSensor := []models.LDRPhotoresistorSensor{}

	fmt.Println(len(res[0].Series))
	if len(res[0].Series) < 1 {
		c.JSON(400, "ID not found data")
		return
	}

	for _, row := range res[0].Series[0].Values {
		LDRPhoto := models.LDRPhotoresistorSensor{}
		LDRPhoto.Time, err = time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		LDRPhoto.Value, _ = strconv.Atoi(string(row[1].(json.Number)))
		LDRPhoto.SensorId, _ = strconv.Atoi(row[2].(string))
		LDRPhoto.StatusAlert, _ = strconv.Atoi(string(row[3].(json.Number)))

		LDRPhotoresistorSensor = append(LDRPhotoresistorSensor, LDRPhoto)
	}

	response := LDRPhotoresistorSensor
	c.JSON(200, response)

}

func CreateLDRPhotoresistorSensor(c *gin.Context) {
	status, valid := functions.IsAuthorized(c.Request.Header, true)
	if status != "Success" && !valid {
		c.JSON(401, status)
		return
	}

	LDRPhotoresistorSensor := models.LDRPhotoresistorSensor{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)

	err := json.Unmarshal(buf.Bytes(), &LDRPhotoresistorSensor)
	if err != nil {
		functions.NotifyToLine("LDRPhotoresistorSensor : Format or data invalid")
		c.String(400, "")
		return
	}

	LDRPhotoresistorSensor.Time = time.Now()
	LDRPhotoresistorSensor.Value = 1023 - LDRPhotoresistorSensor.Value
	if LDRPhotoresistorSensor.Value < 100 {
		LDRPhotoresistorSensor.StatusAlert = 1
	} else if LDRPhotoresistorSensor.Value < 200 {
		LDRPhotoresistorSensor.StatusAlert = 2
	} else if LDRPhotoresistorSensor.Value > 800 {
		LDRPhotoresistorSensor.StatusAlert = 4
	} else if LDRPhotoresistorSensor.Value > 1000 {
		LDRPhotoresistorSensor.StatusAlert = 5
	} else {
		LDRPhotoresistorSensor.StatusAlert = 3
	}

	// Create a point and add to batch
	tags := map[string]string{"sensor_id": fmt.Sprint(LDRPhotoresistorSensor.SensorId)}
	fields := map[string]interface{}{
		"value":        LDRPhotoresistorSensor.Value,
		"status_alert": LDRPhotoresistorSensor.StatusAlert,
	}

	database.Insert(constants.SeriesNameLDRPhotoresistor, tags, fields, LDRPhotoresistorSensor.Time)
	response := LDRPhotoresistorSensor
	c.JSON(200, response)

}
