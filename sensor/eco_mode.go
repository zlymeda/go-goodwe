package sensor

import (
	"bytes"
	"fmt"
)

const (
	sunday    = 1
	monday    = sunday * 2
	tuesday   = monday * 2
	wednesday = tuesday * 2
	thursday  = wednesday * 2
	friday    = thursday * 2
	saturday  = friday * 2
)

type Mode struct {
	Error string
	Start AtTime
	End   AtTime
	Days  Days
	Power int16
	On    bool
}

func (v Mode) IsChargeMode() bool {
	if v.Power > 0 {
		return false
	}

	return v.IsAlwaysOn()
}

func (v Mode) IsDischargeMode() bool {
	if v.Power < 0 {
		return false
	}

	return v.IsAlwaysOn()
}

func (v Mode) IsAlwaysOn() bool {
	alwaysOn := CreateAlwaysOn()
	return v.Start == alwaysOn.Start &&
		v.End == alwaysOn.End &&
		v.On &&
		v.Days == alwaysOn.Days
}

type AtTime struct {
	Hour   int8
	Minute int8
}

type Days struct {
	Mon bool
	Tue bool
	Wed bool
	Thu bool
	Fri bool
	Sat bool
	Sun bool
}

func EcoMode(id string, offset int64, name string) Sensor[Mode] {
	return Sensor[Mode]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   8,
		Unit:   "",
		Kind:   BATTERY,
		read:   ReadEcoModeValue,
		encode: EncodeEcoModeValue,
	}
}

func CreateCharge(power uint) Mode {
	mode := CreateAlwaysOn()
	mode.Power = int16(-1 * int(power))
	return mode
}

func CreateDischarge(power uint) Mode {
	mode := CreateAlwaysOn()
	mode.Power = int16(power)
	return mode
}

func CreateOff() Mode {
	return Mode{
		Start: AtTime{
			Hour:   48,
			Minute: 0,
		},
		End: AtTime{
			Hour:   48,
			Minute: 0,
		},
		Power: 100,
		On:    false,
	}
}

func CreateAlwaysOn() Mode {
	return Mode{
		Error: "",
		Start: AtTime{
			Hour:   0,
			Minute: 0,
		},
		End: AtTime{
			Hour:   23,
			Minute: 59,
		},
		Power: 0,
		Days: Days{
			Mon: true,
			Tue: true,
			Wed: true,
			Thu: true,
			Fri: true,
			Sat: true,
			Sun: true,
		},
		On: true,
	}
}

func ReadEcoModeValue(reader *bytes.Reader) Mode {
	var err error
	mode := Mode{}

	mode.Start, err = readAtTime(reader)
	if err != nil {
		return wrapErr(err)
	}

	mode.End, err = readAtTime(reader)
	if err != nil {
		return wrapErr(err)
	}

	mode.Power, err = readPower(reader)
	if err != nil {
		return wrapErr(err)
	}

	mode.On, err = readOnOff(reader)
	if err != nil {
		return wrapErr(err)
	}

	mode.Days, err = readDays(reader)
	if err != nil {
		return wrapErr(err)
	}

	return mode
}

func EncodeEcoModeValue(value Mode) []byte {
	var result []byte

	result = append(result, uint8(value.Start.Hour))
	result = append(result, uint8(value.Start.Minute))
	result = append(result, uint8(value.End.Hour))
	result = append(result, uint8(value.End.Minute))
	result = append(result, EncodePower2Value(value.Power)...)
	result = append(result, encodeOn(value.On))
	result = append(result, encodeDaysOfWeek(value.Days))

	return result
}

func encodeOn(on bool) byte {
	if on {
		return 0b11111111
	}
	return 0
}

func decodeDaysOfWeek(d uint8) Days {
	return Days{
		Mon: d&monday > 0,
		Tue: d&tuesday > 0,
		Wed: d&wednesday > 0,
		Thu: d&thursday > 0,
		Fri: d&friday > 0,
		Sat: d&saturday > 0,
		Sun: d&sunday > 0,
	}
}

func encodeDaysOfWeek(d Days) uint8 {
	result := uint8(0)

	if d.Mon {
		result = result | monday
	}
	if d.Tue {
		result = result | tuesday
	}
	if d.Wed {
		result = result | wednesday
	}
	if d.Thu {
		result = result | thursday
	}
	if d.Fri {
		result = result | friday
	}
	if d.Sat {
		result = result | saturday
	}
	if d.Sun {
		result = result | sunday
	}

	return result
}

func readOnOff(reader *bytes.Reader) (bool, error) {
	onOff := ReadByteValue(reader)
	if onOff != 0 && onOff != -1 {
		return false, fmt.Errorf("invalid on/off: %d", onOff)
	}
	return onOff != 0, nil
}

func readDays(reader *bytes.Reader) (Days, error) {
	dayBits := readUint8(reader)
	if int8(dayBits) < 0 {
		return Days{}, fmt.Errorf("invalid day bits: %b", dayBits)
	}
	return decodeDaysOfWeek(dayBits), nil
}

func readPower(reader *bytes.Reader) (int16, error) {
	power := ReadPower2Value(reader)
	if power < -100 || power > 100 {
		return 0, fmt.Errorf("invalid power: %d", power)
	}

	return power, nil
}

func maxCharge(reader *bytes.Reader) (int16, error) {
	maxCharge := ReadPower2Value(reader)
	if maxCharge < 0 || maxCharge > 100 {
		return 0, fmt.Errorf("invalid max charge: %d", maxCharge)
	}

	return maxCharge, nil
}

func readAtTime(reader *bytes.Reader) (AtTime, error) {
	hour := readUint8(reader)
	if (hour < 0 || hour > 23) && hour != 48 {
		return AtTime{}, fmt.Errorf("invalid hour: %d", hour)
	}

	minute := readUint8(reader)
	if minute < 0 || minute > 59 {
		return AtTime{}, fmt.Errorf("invalid minute: %d", minute)
	}

	return AtTime{
		Hour:   int8(hour),
		Minute: int8(minute),
	}, nil
}

func wrapErr(err error) Mode {
	return Mode{Error: err.Error()}
}
