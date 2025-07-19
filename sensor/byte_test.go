package sensor_test

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

func TestByte(t *testing.T) {
	tests := []test[int8]{
		{
			offset:   0,
			buffer:   "0c",
			expected: 12,
		},
		{
			offset:   4,
			buffer:   "ff-ee-dd-cc-0c",
			expected: 12,
		},
		{
			offset:   0,
			buffer:   "f0",
			expected: -16,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Byte("", tt.offset, "", "", sensor.PV))
		})
	}
}
