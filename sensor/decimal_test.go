package sensor_test

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

func TestDecimal(t *testing.T) {
	tests := []test[float64]{
		{
			offset:   0,
			buffer:   "00-31",
			expected: 4.9,
		},
		{
			offset:   0,
			buffer:   "ff-9e",
			expected: -9.8,
		},
		{
			offset:   2,
			buffer:   "18-19-00-ff",
			expected: 25.5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Decimal("", tt.offset, 10, "", "", sensor.PV))
		})
	}
}
