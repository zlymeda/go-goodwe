package sensor_test

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

func TestEnergy4(t *testing.T) {
	tests := []test[float64]{
		{
			offset:   0,
			buffer:   "00-02-09-72",
			expected: 13349.0,
		},
		{
			offset:   0,
			buffer:   "ff-ff-ff-ff",
			expected: -0.1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Energy4("", tt.offset, "", sensor.PV))
		})
	}
}

func TestEnergy2(t *testing.T) {
	tests := []test[float64]{
		{
			offset:   0,
			buffer:   "09-72",
			expected: 241.8,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Energy2("", tt.offset, "", sensor.PV))
		})
	}
}
