package sensor_test

import (
	"bytes"
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
)

func TestCalculated(t *testing.T) {
	tests := []test[int16]{
		{
			offset:   0,
			buffer:   "00-31-00-0E",
			expected: 49 + 14,
		},
	}

	s1 := sensor.Integer("", 0, "", "", sensor.PV)
	s2 := sensor.Integer("", 2, "", "", sensor.PV)

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			sensor.Calculated[int16]("", "", func(reader *bytes.Reader) int16 {
				return s1.Read(reader) + s2.Read(reader)
			}, "W", sensor.PV)
		})
	}
}
