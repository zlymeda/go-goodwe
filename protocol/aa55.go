package protocol

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// NewAa55Protocol - Inverter communication protocol seen mostly on older generations of inverters.
//
//	Quite probably it is some variation of the protocol used on RS-485 serial link,
//	extended/adapted to UDP transport layer.
//
//	Each request starts with header of 0xAA, 0x55, then 0xC0, 0x7F (probably some sort of address/command)
//	followed by actual payload data.
//	It is suffixed with 2 bytes of plain checksum of header+payload.
//
//	Response starts again with 0xAA, 0x55, then 0x7F, 0xC0.
//	5-6th bytes are some response type, byte 7 is length of the response payload.
//	The last 2 bytes are again plain checksum of header+payload.
func NewAa55Protocol(payload string, responseType uint16) Request {
	msg, err := encodeAA55Msg(payload)
	if err != nil {
		panic(err)
	}

	return Request{
		Data: msg,
		responseValidator: func(bytes []byte) error {
			return validateAA55Response(bytes, responseType)
		},
	}
}

/*
Validate the response.
data[0:3] is header
data[4:5] is response type
data[6] is response payload length
data[-2:] is checksum (plain sum of response data incl. header)
*/
func validateAA55Response(data []byte, responseType uint16) error {
	if len(data) <= 8 || len(data) != (int(data[6])+9) {
		return fmt.Errorf("invalid response length: %d", len(data))
	}

	if responseType > 0 {
		dataRt := binary.BigEndian.Uint16(data[4:6])
		if dataRt != responseType {
			return fmt.Errorf("invalid response type: %d", dataRt)
		}
	}

	l := len(data)
	checksum := checkSumBytesAA55(data[:(l - 2)])
	dataChecksum := binary.BigEndian.Uint16(data[(l - 2):])
	if uint16(checksum) != dataChecksum {
		return fmt.Errorf("invalid checksum: %d != %d", checksum, dataChecksum)
	}

	return nil
}

func encodeAA55Msg(payload string) ([]byte, error) {
	msg := "AA55C07F" + payload
	checksum := checksumAA55(msg)

	data, err := hex.DecodeString(msg + checksum)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func checksumAA55(msg string) string {
	dst, _ := hex.DecodeString(msg)

	checksum := checkSumBytesAA55(dst)

	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(checksum))

	return hex.EncodeToString(bs)
}

func checkSumBytesAA55(data []byte) int {
	checksum := 0
	for _, x := range data {
		checksum += int(x)
	}
	return checksum
}
