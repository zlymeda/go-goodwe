package sensor

import (
	"bytes"
)

func Frequency(id string, offset int64, name string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   "Hz",
		Kind:   kind,
		read:   ReadFrequencyValue,
		encode: EncodeFrequencyValue,
	}
}

func ReadFrequencyValue(reader *bytes.Reader) float64 {
	return ReadDecimalValue(reader, 100.0)
}

func EncodeFrequencyValue(value float64) []byte {
	return EncodeDecimalValue(value, 100.0)
}
