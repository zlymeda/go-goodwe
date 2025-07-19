package sensor

import (
	"bytes"
	"time"
)

func Timestamp(id string, offset int64, name string) Sensor[time.Time] {
	return Sensor[time.Time]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   6,
		Unit:   "",
		Kind:   NA,
		read:   ReadTimestampValue,
		encode: EncodeTimestampValue,
	}
}

func ReadTimestampValue(reader *bytes.Reader) time.Time {
	year := 2000 + int(readUint8(reader))
	month := time.Month(readUint8(reader))
	day := int(readUint8(reader))
	hour := int(readUint8(reader))
	minute := int(readUint8(reader))
	second := int(readUint8(reader))

	return time.Date(
		year,
		month,
		day,
		hour,
		minute,
		second,
		0,
		time.UTC)
}

func EncodeTimestampValue(value time.Time) []byte {
	return []byte{
		uint8(value.Year() - 2000),
		uint8(value.Month()),
		uint8(value.Day()),
		uint8(value.Hour()),
		uint8(value.Minute()),
		uint8(value.Second()),
	}
}
