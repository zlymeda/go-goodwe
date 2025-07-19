package sensor

import (
	"bytes"
	"golang.org/x/exp/constraints"
)

func Plus[T constraints.Float | constraints.Integer](id, name string, sensor Sensor[T], sensors ...Sensor[T]) Sensor[T] {
	return Calculated(id, name, func(reader *bytes.Reader) T {

		sum := sensor.Read(reader)
		for _, s := range sensors {
			sum += s.Read(reader)
		}

		return sum
	}, sensor.Unit, sensor.Kind)
}

func Mult[T constraints.Float | constraints.Integer](id, name string, unit string, sensor Sensor[T], sensors ...Sensor[T]) Sensor[T] {
	return Calculated(id, name, func(reader *bytes.Reader) T {

		product := sensor.Read(reader)
		for _, s := range sensors {
			product *= s.Read(reader)
		}

		return product
	}, unit, sensor.Kind)
}
