package sensor_test

import (
	"github.com/zlymeda/go-goodwe/internal/assert"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

type testEcoMode struct {
	name string
	test[sensor.Mode]
}

func TestEcoMode(t *testing.T) {
	tests := []testEcoMode{
		{
			name: "13:30-14:40 Mon,Wed,Thu -60% On",
			test: test[sensor.Mode]{
				offset: 0,
				buffer: "0d-1e-0e-28-ff-c4-ff-1a",
				expected: sensor.Mode{
					Error: "",
					Start: sensor.AtTime{
						Hour:   13,
						Minute: 30,
					},
					End: sensor.AtTime{
						Hour:   14,
						Minute: 40,
					},
					Power: -60,
					Days: sensor.Days{
						Mon: true,
						Wed: true,
						Thu: true,
					},
					On: true,
				},
			},
		},
		{
			name: "Off",
			test: test[sensor.Mode]{
				offset:   0,
				buffer:   "30-00-30-00-00-64-00-00",
				expected: sensor.CreateOff(),
			},
		},
		{
			name: "Charge at 40",
			test: test[sensor.Mode]{
				offset:   0,
				buffer:   "00-00-17-3b-ff-d8-ff-7f",
				expected: sensor.CreateCharge(40),
			},
		},
		{
			name: "Discharge at 60",
			test: test[sensor.Mode]{
				offset:   0,
				buffer:   "00-00-17-3b-00-3c-ff-7f",
				expected: sensor.CreateDischarge(60),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertConversion(t, tt.test, sensor.EcoMode("", tt.offset, ""))
		})
	}
}

func TestEcoModeChargeAndDischargeModes(t *testing.T) {

	dischargeMode := sensor.CreateDischarge(50)

	assert.Equals(t, true, dischargeMode.IsAlwaysOn())
	assert.Equals(t, true, dischargeMode.IsDischargeMode())
	assert.Equals(t, false, dischargeMode.IsChargeMode())

	chargeMode := sensor.CreateCharge(50)

	assert.Equals(t, true, chargeMode.IsAlwaysOn())
	assert.Equals(t, false, chargeMode.IsDischargeMode())
	assert.Equals(t, true, chargeMode.IsChargeMode())

}
