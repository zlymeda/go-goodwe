package protocol

import (
	"github.com/zlymeda/go-goodwe/internal/assert"
	"testing"
)

func TestChecksumAA55(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "0123456789", expected: "0159"},
		{input: "00", expected: "0000"},
		{input: "0AFF", expected: "0109"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			output := checksumAA55(tt.input)
			assert.Equals(t, tt.expected, output)

		})
	}
}
