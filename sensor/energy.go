package sensor

import (
	"bytes"
)

func Energy2(id string, offset int64, name string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   "kWh",
		Kind:   kind,
		read:   ReadEnergyValue,
		encode: EncodeEnergyValue,
	}
}

func ReadEnergyValue(reader *bytes.Reader) float64 {
	return ReadDecimalValue(reader, 10)
}

func EncodeEnergyValue(value float64) []byte {
	return EncodeDecimalValue(value, 10)
}

func Energy4(id string, offset int64, name string, kind Kind) Sensor[float64] {
	return Sensor[float64]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   4,
		Unit:   "kWh",
		Kind:   kind,
		read:   ReadEnergy4Value,
		encode: EncodeEnergy4Value,
	}
}

func ReadEnergy4Value(reader *bytes.Reader) float64 {
	return ReadFloatValue(reader, 10.0)
}

func EncodeEnergy4Value(value float64) []byte {
	return EncodeFloatValue(value, 10)
}
