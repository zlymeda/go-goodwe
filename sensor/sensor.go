package sensor

import (
	"bytes"
	"io"
)

type Sizeable interface {
	GetOffset() int64
	GetSize() int
}

type BaseSensor interface {
	GetId() string
	GetName() string
	GetKind() Kind
	GetUnit() string
}

type reader[T any] func(reader *bytes.Reader) T
type encoder[T any] func(value T) []byte

type Sensor[T any] struct {
	Id     string
	Offset int64
	Name   string
	Size   int
	Unit   string
	Kind   Kind

	read   reader[T]
	encode encoder[T]
}

func (s Sensor[T]) GetId() string {
	return s.Id
}

func (s Sensor[T]) GetName() string {
	return s.Name
}

func (s Sensor[T]) GetKind() Kind {
	return s.Kind
}

func (s Sensor[T]) GetUnit() string {
	return s.Unit
}

func (s Sensor[T]) Read(reader *bytes.Reader) T {
	_, err := reader.Seek(s.Offset, io.SeekStart)
	if err != nil {
		panic(err)
	}

	return s.read(reader)
}

func (s Sensor[T]) JustRead(reader *bytes.Reader) T {
	return s.read(reader)
}

func (s Sensor[T]) Decode(input T) []byte {
	return s.encode(input)
}

func (s Sensor[T]) GetOffset() int64 {
	return s.Offset
}

func (s Sensor[T]) GetSize() int {
	return s.Size
}

func (s Sensor[T]) ReadUsing(read func(s Sizeable) *bytes.Reader) T {
	return s.JustRead(read(s))
}
