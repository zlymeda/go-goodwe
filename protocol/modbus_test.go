package protocol

import (
	"encoding/hex"
	"github.com/zlymeda/go-goodwe/internal/assert"
	"testing"
)

func TestCreateModBus(t *testing.T) {
	tests := []struct {
		name     string
		input    Request
		expected string
	}{
		{
			input:    NewModbusRead(0xf7, 0x88b8, 0x0021),
			expected: "f70388b800213ac1",
		},
		{
			input:    NewModbusWrite(0xf7, 0x88b8, 0x00FF),
			expected: "f70688b800ff7699",
		},
		{
			input:    NewModbusWriteMulti(0xf7, 0x88b8, []byte("\x01\x02\x03\x04\x05\x06")),
			expected: "f71088b8000306010203040506102e",
		},
		{
			input:    NewModbusRead(0xf7, 0x88b8, 0x0021),
			expected: "f70388b800213ac1",
		},
		{
			input:    NewModbusWrite(0xf7, 0xb798, 0x0002),
			expected: "f706b7980002bac6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			output := hex.EncodeToString(tt.input.Data)
			assert.Equals(t, tt.expected, output)
		})
	}
}

func TestValidateModBusReadResponse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "aa55f7030401020304cd33",
			expected: "",
		},
		{
			input:    "aa55f7030401020304cd33ffffff",
			expected: "",
		},
		{
			input:    "aa55f7030401020304",
			expected: "read; response is too short: 9, expected 11",
		},
		{
			input:    "aa55f70304010203043346",
			expected: "response CRC-16 checksum does not match",
		},
		{
			input:    "aa55f783040102030405b35e",
			expected: "response is command failure: SLAVE DEVICE FAILURE",
		},
		{
			input:    "aa55f70306010203040506b417",
			expected: "read: response has unexpected length: 6, expected 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			validateResponse(t, tt.input, tt.expected, 0x03, 0x0401, 2)
		})
	}
}

func TestValidateModBusWriteResponse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "aa55f706b12c00147ba6",
			expected: "",
		},
		{
			input:    "aa55f706b12c00147ba6ffffff",
			expected: "",
		},
		{
			input:    "aa55f7066b12",
			expected: "write: response has unexpected length: 6, expected 10",
		},
		{
			input:    "aa55f706b12c00147ba7",
			expected: "response CRC-16 checksum does not match",
		},
		{
			input:    "aa55f706b12c0012fba4",
			expected: "write: response has wrong value: 12, expected 14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			validateResponse(t, tt.input, tt.expected, 0x06, 0xb12c, 0x14)
		})
	}
}

func TestValidateModBusWriteMultiResponse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "aa55f71088b800033f1b",
			expected: "",
		},
		{
			input:    "aa55f71088b800033f1bffffff",
			expected: "",
		},
		{
			input:    "aa55f71088b8",
			expected: "write: response has unexpected length: 6, expected 10",
		},
		{
			input:    "aa55f71088b800033f1c",
			expected: "response CRC-16 checksum does not match",
		},
		{
			input:    "aa55f71088b80001beda",
			expected: "write: response has wrong value: 1, expected 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			validateResponse(t, tt.input, tt.expected, 0x10, 0x88b8, 0x03)
		})
	}
}

func validateResponse(t *testing.T, input string, expected string, cmd uint8, offset uint16, value uint16) {
	decoded, err := hex.DecodeString(input)
	assert.Ok(t, err)

	result := ""
	errResult := validateModbusResponse(decoded, cmd, offset, value)
	if errResult != nil {
		result = errResult.Error()
	}

	assert.Equals(t, expected, result)
}
