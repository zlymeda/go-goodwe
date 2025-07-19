package sensor

import (
	"bytes"
)

func Float(id string, offset int64, scale int, name string, unit string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   4,
		Unit:   unit,
		Kind:   kind,
		read: func(reader *bytes.Reader) float64 {
			return ReadFloatValue(reader, scale)
		},
		encode: func(value float64) []byte {
			return EncodeFloatValue(value, scale)
		},
	}
}

func ReadFloatValue(reader *bytes.Reader, scale int) float64 {
	intVal := ReadLongValue(reader)
	return float64(intVal) / float64(scale)
}

func EncodeFloatValue(value float64, scale int) []byte {
	intVal := int32(value * float64(scale))
	return EncodeLongValue(intVal)
}
