package sensor_test

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

func TestPower4(t *testing.T) {
	tests := []test[int32]{
		{
			offset:   0,
			buffer:   "00-00-06-9f",
			expected: 1695,
		},
		{
			offset:   0,
			buffer:   "ff-ff-ff-fd",
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Power4("", tt.offset, "", sensor.PV))
		})
	}
}

func TestPower2(t *testing.T) {
	tests := []test[int16]{
		{
			offset:   0,
			buffer:   "06-9f",
			expected: 1695,
		},
		{
			offset:   0,
			buffer:   "ff-fd",
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Power2("", tt.offset, "", sensor.PV))
		})
	}
}
