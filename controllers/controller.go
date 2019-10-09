package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MySmartFarm/mysmartfarm_api/constants"
	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	//Parameter from url
	fmt.Println("Hello")
	//Return json
	c.JSON(http.StatusOK, "")
}

func GetENV(c *gin.Context) {
	type EnvShow struct {
		MSFEnvironmentModeKey string `json:"MSFEnvironmentModeKey"`
		InfluxdbUsername      string `json:"InfluxdbUsername"`
		InfluxdbPassword      string `json:"InfluxdbPassword"`
		NotifyLineToken       string `json:"NotifyLineToken"`
	}
	envShow := EnvShow{}
	envShow.MSFEnvironmentModeKey = os.Getenv(constants.MSFEnvironmentModeKey)
	envShow.InfluxdbUsername = os.Getenv(constants.InfluxdbUsername)
	envShow.InfluxdbPassword = os.Getenv(constants.InfluxdbPassword)
	envShow.NotifyLineToken = os.Getenv(constants.NotifyLineToken)
	c.JSON(http.StatusOK, envShow)
}
