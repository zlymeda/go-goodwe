package sensor

import (
	"bytes"
)

func Decimal(id string, offset int64, scale int, name string, unit string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   unit,
		Kind:   kind,
		read: func(reader *bytes.Reader) float64 {
			return ReadDecimalValue(reader, scale)
		},
		encode: func(value float64) []byte {
			return EncodeDecimalValue(value, scale)
		},
	}
}

func ReadDecimalValue(reader *bytes.Reader, scale int) float64 {
	intVal := ReadIntegerValue(reader)
	return float64(intVal) / float64(scale)
}

func EncodeDecimalValue(value float64, scale int) []byte {
	intVal := int16(value * float64(scale))
	return EncodeIntegerValue(intVal)
}
