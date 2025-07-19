package sensor

import (
	"bytes"
)

func Current(id string, offset int64, name string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   "A",
		Kind:   kind,
		read:   ReadCurrentValue,
		encode: EncodeCurrentValue,
	}
}

func ReadCurrentValue(reader *bytes.Reader) float64 {
	return ReadDecimalValue(reader, 10.0)
}

func EncodeCurrentValue(value float64) []byte {
	return EncodeDecimalValue(value, 10.0)
}
