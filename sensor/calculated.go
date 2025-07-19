package sensor

import "bytes"

func Calculated[T any](id string, name string, calculate func(reader *bytes.Reader) T, unit string, kind Kind) Sensor[T] {
	return Sensor[T]{
		Id:     id,
		Offset: 0,
		Name:   name,
		Size:   0,
		Unit:   unit,
		Kind:   kind,
		read:   calculate,
		encode: func(value T) []byte {
			return nil
		},
	}
}
