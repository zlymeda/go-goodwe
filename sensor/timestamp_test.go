package sensor_test

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/sensor"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	tests := []test[time.Time]{
		{
			offset:   0,
			buffer:   "16-01-04-12-1e-19",
			expected: time.Date(2022, 1, 4, 18, 30, 25, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s at %d", tt.buffer, tt.offset), func(t *testing.T) {
			assertConversion(t, tt, sensor.Timestamp("", tt.offset, ""))
		})
	}
}
