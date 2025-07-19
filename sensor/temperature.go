package sensor

import (
	"bytes"
)

func Temperature(id string, offset int64, name string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   "°C",
		Kind:   kind,
		read:   ReadTemperatureValue,
		encode: EncodeTemperatureValue,
	}
}

func ReadTemperatureValue(reader *bytes.Reader) float64 {
	return ReadDecimalValue(reader, 10.0)
}

func EncodeTemperatureValue(value float64) []byte {
	return EncodeDecimalValue(value, 10.0)
}
