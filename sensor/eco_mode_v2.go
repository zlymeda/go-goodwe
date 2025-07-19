package sensor

import (
	"bytes"
)

type ModeV2 struct {
	Mode
	MaxCharge int16
}

func (v ModeV2) IsChargeMode() bool {
	return v.Mode.IsChargeMode()
}

func (v ModeV2) IsDischargeMode() bool {
	return v.Mode.IsDischargeMode()
}

func EcoModeV2(id string, offset int64, name string) Sensor[ModeV2] {
	return Sensor[ModeV2]{
		Id:     id,
		Offset: offset,
		Name:   name,
		Size:   12,
		Unit:   "",
		Kind:   BATTERY,
		read:   ReadEcoModeV2Value,
		encode: EncodeEcoModeV2Value,
	}
}

func CreateChargeV2(power, maxCharge uint) ModeV2 {
	return ModeV2{
		MaxCharge: int16(maxCharge),
		Mode:      CreateCharge(power),
	}
}

func CreateDischargeV2(power uint) ModeV2 {
	return ModeV2{
		MaxCharge: 100,
		Mode:      CreateDischarge(power),
	}
}

func CreateOffV2() ModeV2 {
	return ModeV2{
		MaxCharge: 100,
		Mode:      CreateOff(),
	}
}

func ReadEcoModeV2Value(reader *bytes.Reader) ModeV2 {
	var err error
	mode := ModeV2{}

	mode.Start, err = readAtTime(reader)
	if err != nil {
		return wrapErrV2(err)
	}

	mode.End, err = readAtTime(reader)
	if err != nil {
		return wrapErrV2(err)
	}

	mode.On, err = readOnOff(reader)
	if err != nil {
		return wrapErrV2(err)
	}

	mode.Days, err = readDays(reader)
	if err != nil {
		return wrapErrV2(err)
	}

	mode.Power, err = readPower(reader)
	if err != nil {
		return wrapErrV2(err)
	}

	mode.MaxCharge, err = maxCharge(reader)
	if err != nil {
		return wrapErrV2(err)
	}

	return mode
}

func EncodeEcoModeV2Value(value ModeV2) []byte {
	var result []byte

	result = append(result, uint8(value.Start.Hour))
	result = append(result, uint8(value.Start.Minute))
	result = append(result, uint8(value.End.Hour))
	result = append(result, uint8(value.End.Minute))
	result = append(result, encodeOn(value.On))
	result = append(result, encodeDaysOfWeek(value.Days))
	result = append(result, EncodePower2Value(value.Power)...)
	result = append(result, EncodePower2Value(value.MaxCharge)...)
	result = append(result, 0)
	result = append(result, 0)

	return result
}

func wrapErrV2(err error) ModeV2 {
	return ModeV2{
		Mode: Mode{Error: err.Error()},
	}
}
