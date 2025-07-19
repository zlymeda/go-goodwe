package protocol

import (
	"encoding/binary"
	"fmt"
	"log/slog"
)

const (
	ModbusReadCmd       = 0x3
	ModbusWriteCmd      = 0x6
	ModbusWriteMultiCmd = 0x10
)

var (
	failureCodes = map[uint8]string{
		1:  "ILLEGAL FUNCTION",
		2:  "ILLEGAL DATA ADDRESS",
		3:  "ILLEGAL DATA VALUE",
		4:  "SLAVE DEVICE FAILURE",
		5:  "ACKNOWLEDGE",
		6:  "SLAVE DEVICE BUSY",
		7:  "NEGATIVE ACKNOWLEDGEMENT",
		8:  "MEMORY PARITY ERROR",
		10: "GATEWAY PATH UNAVAILABLE",
		11: "GATEWAY TARGET DEVICE FAILED TO RESPOND",
	}
	crc16table = createCrc16Table()
)

func NewModbusRead(commAddr uint8, offset uint16, count uint16) Request {
	return newModbus(commAddr, ModbusReadCmd, offset, count)
}

func NewModbusWrite(commAddr uint8, register uint16, value uint16) Request {
	return newModbus(commAddr, ModbusWriteCmd, register, value)
}

func NewModbusWriteMulti(commAddr uint8, offset uint16, values []byte) Request {
	return newModbusMulti(commAddr, ModbusWriteMultiCmd, offset, values)
}

// newModbus creates a representation of a modbus request:
// Inverter communication protocol seen on newer generation of inverters, based on Modbus
// protocol over UDP transport layer.
// The modbus communication is rather simple, there are "registers" at specified addresses/offsets,
// each represented by 2 bytes. The protocol may query/update individual or range of these registers.
// Each register represents some measured value or operational settings.
// It's inverter implementation specific which register means what.
// Some values may span more registers (i.e. 4bytes measurement value over 2 registers).
//
// Every request usually starts with communication address (usually 0xF7, but can be changed).
// Second byte is the modbus command - 0x03 read multiple, 0x06 write single, 0x10 write multiple.
// Bytes 3-4 represent the register address (or start of range)
// Bytes 5-6 represent the command parameter (range size or actual value for write).
// Last 2 bytes of request is the CRC-16 (modbus flavor) of the request.
//
// Responses seem to always start with 0xAA, 0x55, then the comm_addr and modbus command.
// (If the command fails, the highest bit of command is set to 1 ?)
// For read requests, next byte is response payload length, then the actual payload.
// Last 2 bytes of response is again the CRC-16 of the response.
func newModbus(commAddr uint8, cmd uint8, offset uint16, value uint16) Request {
	return newModbusRequest(
		createModbusRequest(commAddr, cmd, offset, value), cmd, offset, value)
}

func newModbusMulti(commAddr uint8, cmd uint8, offset uint16, values []byte) Request {
	return newModbusRequest(
		createModbusMultiRequest(commAddr, cmd, offset, values), cmd, offset, uint16(len(values))/2)
}

// newModbus creates modbus request
//
//	data[0] is inverter address
//	data[1] is modbus command
//	data[2:3] is command offset parameter
//	data[4:5] is command value parameter
//	data[6:7] is crc-16 checksum
func newModbusRequest(data []byte, cmd uint8, offset uint16, value uint16) Request {
	return Request{
		Data: data,
		responseValidator: func(bytes []byte) error {
			err := validateModbusResponse(bytes, cmd, offset, value)
			if err != nil {
				slog.Warn("newModbusRequest", slog.Any("err", err))
				// log.Printf("WARN: %+v\n", err)
			}
			return err
		},
	}
}

func createModbusRequest(commAddr uint8, cmd uint8, offset uint16, value uint16) []byte {
	data := []byte{
		commAddr,
		cmd,
		uint8((offset >> 8) & 0xFF),
		uint8(offset & 0xFF),
		uint8((value >> 8) & 0xFF),
		uint8(value & 0xFF),
	}

	checksum := modbusChecksum(data)

	data = append(data, uint8(checksum&0xFF))
	data = append(data, uint8((checksum>>8)&0xFF))
	return data
}

func createModbusMultiRequest(commAddr uint8, cmd uint8, offset uint16, values []byte) []byte {
	data := []byte{
		commAddr,
		cmd,
		uint8((offset >> 8) & 0xFF),
		uint8(offset & 0xFF),
		0,
		uint8(len(values) / 2),
		uint8(len(values)),
	}

	data = append(data, values...)

	checksum := modbusChecksum(data)

	data = append(data, uint8(checksum&0xFF))
	data = append(data, uint8((checksum>>8)&0xFF))
	return data
}

// validateModbusResponse
//
//	data[0:1] is header
//	data[2] is source address
//	data[3] is command return type
//	data[4] is response payload length (for read commands)
//	data[-2:] is crc-16 checksum
func validateModbusResponse(data []byte, cmd uint8, offset uint16, value uint16) error {
	if len(data) <= 4 {
		return fmt.Errorf("response is too short")
	}

	if data[3] != cmd {
		failure := failureCodes[data[4]]
		if failure == "" {
			failure = "UNKNOWN"
		}
		return fmt.Errorf("response is command failure: %s", failure)
	}

	var expectedLength int
	dataCmd := data[3]
	if dataCmd == ModbusReadCmd {
		if data[4] != uint8(value*2) {
			return fmt.Errorf("read: response has unexpected length: %d, expected %d", data[4], value*2)
		}

		expectedLength = int(data[4]) + 7
		if len(data) < expectedLength {
			return &PartialResponseErr{
				Message: fmt.Sprintf("read; response is too short: %d, expected %d", len(data), expectedLength),
			}
		}

	} else if dataCmd == ModbusWriteCmd || dataCmd == ModbusWriteMultiCmd {
		if len(data) < 10 {
			return fmt.Errorf("write: response has unexpected length: %d, expected %d", len(data), 10)
		}

		expectedLength = 10

		responseOffset := binary.BigEndian.Uint16(data[4:6])
		if responseOffset != offset {
			return fmt.Errorf("write: response has wrong offset: %X, expected %X", responseOffset, offset)
		}

		responseValue := binary.BigEndian.Uint16(data[6:8])
		if responseValue != value {
			return fmt.Errorf("write: response has wrong value: %X, expected %X", responseValue, value)
		}

	} else {
		expectedLength = len(data)
	}

	checksumOffset := expectedLength - 2
	if checksumOffset <= 0 {
		return fmt.Errorf("invalid checksum offset")
	}

	if modbusChecksum(data[2:checksumOffset]) != (uint16(data[checksumOffset+1])<<8)+uint16(data[checksumOffset]) {
		return fmt.Errorf("response CRC-16 checksum does not match")
	}

	if data[3] != cmd {
		failureCode, ok := failureCodes[4]
		if !ok {
			failureCode = "UNKNOWN"
		}
		return fmt.Errorf("request rejected: %s", failureCode)
	}

	return nil
}

func createCrc16Table() []uint16 {
	var table []uint16

	for i := 0; i < 256; i++ {
		buffer := uint16(i << 1)
		crc := uint16(0)

		for i := 8; i > 0; i-- {
			buffer >>= 1
			if (buffer^crc)&0x0001 != 0 {
				crc = (crc >> 1) ^ 0xA001
			} else {
				crc >>= 1
			}
		}

		table = append(table, crc)
	}

	return table
}

func modbusChecksum(bytes []byte) uint16 {
	crc := uint16(0xFFFF)

	for _, ch := range bytes {
		crc = (crc >> 8) ^ crc16table[(crc^uint16(ch))&0xFF]
	}

	return crc
}
