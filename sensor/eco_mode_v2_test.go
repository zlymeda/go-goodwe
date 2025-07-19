package sensor_test

import (
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

type testEcoModeV2 struct {
	name string
	test[sensor.ModeV2]
}

func TestEcoModeV2(t *testing.T) {
	tests := []testEcoModeV2{
		{
			name: "13:30-14:40 Mon,Wed,Thu -60% (max charge 90%) On",
			test: test[sensor.ModeV2]{
				offset: 0,
				buffer: "0d-1e-0e-28-ff-1a-ff-c4-00-5a-00-00",
				expected: sensor.ModeV2{
					MaxCharge: 90,
					Mode: sensor.Mode{
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
		},
		{
			name: "Off",
			test: test[sensor.ModeV2]{
				offset:   0,
				buffer:   "30-00-30-00-00-00-00-64-00-64-00-00",
				expected: sensor.CreateOffV2(),
			},
		},
		{
			name: "Charge at 40 and 80 max charge",
			test: test[sensor.ModeV2]{
				offset:   0,
				buffer:   "00-00-17-3b-ff-7f-ff-d8-00-50-00-00",
				expected: sensor.CreateChargeV2(40, 80),
			},
		},
		{
			name: "Discharge at 60",
			test: test[sensor.ModeV2]{
				offset:   0,
				buffer:   "00-00-17-3b-ff-7f-00-3c-00-64-00-00",
				expected: sensor.CreateDischargeV2(60),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertConversion(t, tt.test, sensor.EcoModeV2("", tt.offset, ""))
		})
	}
}
