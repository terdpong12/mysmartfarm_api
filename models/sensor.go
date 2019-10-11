package models

import (
	"time"
)

type SoilMoistureSensor struct {
	Time        time.Time `json:"time"`
	SensorId    int       `json:"sensor_id"`
	Value       int       `json:"value"`
	StatusAlert int       `json:"status_alert"`
}

type LDRPhotoresistorSensor struct {
	Time        time.Time `json:"time"`
	SensorId    int       `json:"sensor_id"`
	Value       int       `json:"value"`
	StatusAlert int       `json:"status_alert"`
}
