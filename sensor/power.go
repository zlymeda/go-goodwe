package sensor

import (
	"bytes"
)

func Power2(id string, offset int64, name string, kind Kind) Sensor[int16] {
	return Sensor[int16]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   "W",
		Kind:   kind,
		read:   ReadPower2Value,
		encode: EncodePower2Value,
	}
}

func ReadPower2Value(reader *bytes.Reader) int16 {
	return ReadIntegerValue(reader)
}

func EncodePower2Value(value int16) []byte {
	return EncodeIntegerValue(value)
}

func Power4(id string, offset int64, name string, kind Kind) Sensor[int32] {
	return Sensor[int32]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   4,
		Unit:   "W",
		Kind:   kind,
		read:   ReadPower4Value,
		encode: EncodePower4Value,
	}
}

func ReadPower4Value(reader *bytes.Reader) int32 {
	return ReadLongValue(reader)
}

func EncodePower4Value(value int32) []byte {
	return EncodeLongValue(value)
}
