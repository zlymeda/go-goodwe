package inverter

import (
	"time"
)

type Event struct {
	Time    time.Time
	Runtime ETSensorValues
	Meter   ETSensorsMeterValues
	Battery ETSensorsBatteryValues
}
