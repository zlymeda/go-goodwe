package sensor_test

import (
	"bytes"
	"encoding/hex"
	"github.com/zlymeda/go-goodwe/internal/assert"
	"github.com/zlymeda/go-goodwe/sensor"
	"strings"
	"testing"
)

type test[T any] struct {
	offset   int64
	buffer   string
	expected T
}

func assertConversion[T any](t *testing.T, tt test[T], underTest sensor.Sensor[T]) {
	t.Helper()

	bytez := bytes.NewReader(decode(t, tt.buffer))
	assert.Equals(t, tt.expected, underTest.Read(bytez))

	if tt.offset == 0 {
		assert.Equals(t, unprettify(tt.buffer), hex.EncodeToString(underTest.Decode(tt.expected)))
	}
}

func unprettify(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

func decode(t *testing.T, s string) []byte {
	d, err := hex.DecodeString(unprettify(s))
	assert.Ok(t, err)
	return d
}
