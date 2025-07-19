package sensor

import (
	"bytes"
)

func Integer(id string, offset int64, name string, unit string, kind Kind) Sensor[int16] {
	return Sensor[int16]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   unit,
		Kind:   kind,
		read:   ReadIntegerValue,
		encode: EncodeIntegerValue,
	}
}

func ReadIntegerValue(reader *bytes.Reader) int16 {
	u := readUint16(reader)
	return int16(u)
}

func EncodeIntegerValue(value int16) []byte {
	return writeUint16(uint16(value))
}
