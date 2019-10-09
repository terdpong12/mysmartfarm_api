package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MySmartFarm/mysmartfarm_api/constants"
	"github.com/MySmartFarm/mysmartfarm_api/functions"
	client "github.com/influxdata/influxdb1-client/v2"
)

// MyDB specifies name of database
var URL = "http://localhost:8086"
var MyDB = "mysmartfarm"
var username = "admin"
var password = "secret"

// queryDB convenience function to query the database
func Query(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	c := influxDBClient()
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// Write a point using the HTTP client
func Insert(seriesName string, tags map[string]string, fields map[string]interface{}, timestemp time.Time) {
	// Make client
	c := influxDBClient()
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	pt, err := client.NewPoint(seriesName, tags, fields, timestemp)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)

}

// Create a Database with a query
func CreateDatabase() {
	// Make client
	c := influxDBClient()
	defer c.Close()

	q := client.NewQuery("CREATE DATABASE "+MyDB, "", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}

func GetAddr() string {
	mode := os.Getenv(constants.MSFEnvironmentModeKey)
	switch mode {
	case "Local":
		URL = os.Getenv(constants.InfluxdbURL)
		username = os.Getenv(constants.InfluxdbUsername)
		password = os.Getenv(constants.InfluxdbPassword)
		return URL
	case "Prod":
		URL = os.Getenv(constants.InfluxdbURL)
		username = os.Getenv(constants.InfluxdbUsername)
		password = os.Getenv(constants.InfluxdbPassword)
		return URL
	default:
		os.Exit(1)
	}
	return ""
}

func influxDBClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     GetAddr(),
		Username: username,
		Password: password,
	})
	if err != nil {
		functions.NotifyToLine(fmt.Sprint(err))
		log.Fatalln("Error: ", err)
	}
	return c
}
