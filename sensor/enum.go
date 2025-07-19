package sensor

import (
	"bytes"
)

func Enum(id string, offset int64, name string, labels map[uint8]string, kind Kind) Sensor[string] {
	return Sensor[string]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   1,
		Unit:   "enum",
		Kind:   kind,
		read: func(reader *bytes.Reader) string {
			return ReadEnumValue(reader, labels)
		},
		encode: nil,
	}
}

func ReadEnumValue(reader *bytes.Reader, labels map[uint8]string) string {
	return labels[readUint8(reader)]
}

func Enum2(id string, offset int64, name string, labels map[uint16]string, kind Kind) Sensor[string] {
	return Sensor[string]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   2,
		Unit:   "enum",
		Kind:   kind,
		read: func(reader *bytes.Reader) string {
			return ReadEnum2Value(reader, labels)
		},
		encode: nil,
	}
}

func ReadEnum2Value(reader *bytes.Reader, labels map[uint16]string) string {
	return labels[readUint16(reader)]
}
