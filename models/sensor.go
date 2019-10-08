package models

import (
	"time"
)

type SoilMoistureSensor struct {
	Time          time.Time `json:"time"`
	SensorId      int       `json:"sensor_id"`
	MoistureValue int       `json:"moisture_value"`
	StatusAlert   int       `json:"status_alert"`
}
