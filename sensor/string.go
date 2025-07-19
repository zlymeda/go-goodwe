package sensor

import (
	"bytes"
	"io"
	"strings"
)

func String(id string, name string, offset, end int64) Sensor[string] {
	return Sensor[string]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   1,
		Unit:   "string",
		Kind:   NA,
		read: func(reader *bytes.Reader) string {
			data := make([]byte, end-offset)
			if _, err := io.ReadFull(reader, data); err != nil {
				return err.Error()
			}
			return strings.TrimSpace(string(data))
		},
		encode: nil,
	}
}
