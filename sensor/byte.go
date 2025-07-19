package sensor

import (
	"bytes"
)

func Byte(id string, offset int64, name string, unit string, kind Kind) Sensor[int8] {
	return Sensor[int8]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   1,
		Unit:   unit,
		Kind:   kind,
		read:   ReadByteValue,
		encode: EncodeByteValue,
	}
}

func ReadByteValue(reader *bytes.Reader) int8 {
	u := readUint8(reader)
	return int8(u)
}

func EncodeByteValue(value int8) []byte {
	return []byte{uint8(value)}
}
