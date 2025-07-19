package sensor_test

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

func TestInteger(t *testing.T) {
	tests := []test[int16]{
		{
			offset:   0,
			buffer:   "00-31",
			expected: 49,
		},
		{
			offset:   0,
			buffer:   "ff-9e",
			expected: -98,
		},
		{
			offset:   2,
			buffer:   "18-19-00-ff",
			expected: 255,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Integer("", tt.offset, "", "", sensor.PV))
		})
	}
}
