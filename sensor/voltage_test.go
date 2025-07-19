package sensor_test

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

func TestVoltage(t *testing.T) {
	tests := []test[float64]{
		{
			offset:   0,
			buffer:   "0c-fe",
			expected: 332.6,
		},
		{
			offset:   0,
			buffer:   "1f-64",
			expected: 803.6,
		},
		{
			offset:   0,
			buffer:   "ff-64",
			expected: -15.6,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Voltage("", tt.offset, "", sensor.PV))
		})
	}
}
