package sensor

import (
	"bytes"
)

func Long(id string, offset int64, name string, unit string, kind Kind) Sensor[int32] {
	return Sensor[int32]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   4,
		Unit:   unit,
		Kind:   kind,
		read:   ReadLongValue,
		encode: EncodeLongValue,
	}
}

func ReadLongValue(reader *bytes.Reader) int32 {
	u := readUint32(reader)
	return int32(u)
}

func EncodeLongValue(value int32) []byte {
	return writeUint32(uint32(value))
}
