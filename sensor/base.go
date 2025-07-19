package sensor

import (
	"bytes"
	"encoding/binary"
)

func readUint8(reader *bytes.Reader) uint8 {
	b, _ := reader.ReadByte()
	return b
}

func writeUint16(val uint16) []byte {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, val)
	return bs
}

func readUint16(reader *bytes.Reader) uint16 {
	bs := make([]byte, 0, 2)

	for i := 0; i < 2; i++ {
		b, _ := reader.ReadByte()
		bs = append(bs, b)
	}

	return binary.BigEndian.Uint16(bs)
}

func writeUint32(val uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, val)
	return bs
}

func readUint32(reader *bytes.Reader) uint32 {
	bs := make([]byte, 0, 4)

	for i := 0; i < 4; i++ {
		b, _ := reader.ReadByte()
		bs = append(bs, b)
	}

	return binary.BigEndian.Uint32(bs)
}
