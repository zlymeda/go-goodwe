package sensor

import (
	"bytes"
)

func Voltage(id string, offset int64, name string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   "V",
		Kind:   kind,
		read:   ReadVoltageValue,
		encode: EncodeVoltageValue,
	}
}

func ReadVoltageValue(reader *bytes.Reader) float64 {
	return ReadDecimalValue(reader, 10.0)
}

func EncodeVoltageValue(value float64) []byte {
	return EncodeDecimalValue(value, 10.0)
}
