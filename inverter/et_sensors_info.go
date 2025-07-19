package inverter

import (
	"bytes"
	"github.com/zlymeda/go-goodwe/sensor"
)

//goland:noinspection SpellCheckingInspection
type ETInfo struct {
	ModbusVersion   sensor.Sensor[int16]
	RatedPower      sensor.Sensor[int16]
	AcOutputType    sensor.Sensor[int16]
	SerialNumber    sensor.Sensor[string]
	ModelName       sensor.Sensor[string]
	Dsp1SwVersion   sensor.Sensor[int16]
	Dsp2SwVersion   sensor.Sensor[int16]
	DspSvnVersion   sensor.Sensor[int16]
	ArmSwVersion    sensor.Sensor[int16]
	ArmSvnVersion   sensor.Sensor[int16]
	SoftwareVersion sensor.Sensor[string]
	ArmVersion      sensor.Sensor[string]
}

func (s *ETInfo) Read(reader *bytes.Reader) ETInfoValues {
	if reader == nil {
		return ETInfoValues{}
	}

	return ETInfoValues{
		ModbusVersion:   int(s.ModbusVersion.Read(reader)),
		RatedPower:      int(s.RatedPower.Read(reader)),
		AcOutputType:    int(s.AcOutputType.Read(reader)),
		SerialNumber:    s.SerialNumber.Read(reader),
		ModelName:       s.ModelName.Read(reader),
		Dsp1SwVersion:   int(s.Dsp1SwVersion.Read(reader)),
		Dsp2SwVersion:   int(s.Dsp2SwVersion.Read(reader)),
		DspSvnVersion:   int(s.DspSvnVersion.Read(reader)),
		ArmSwVersion:    int(s.ArmSwVersion.Read(reader)),
		ArmSvnVersion:   int(s.ArmSvnVersion.Read(reader)),
		SoftwareVersion: s.SoftwareVersion.Read(reader),
		ArmVersion:      s.ArmVersion.Read(reader),
	}
}

//goland:noinspection SpellCheckingInspection
type ETInfoValues struct {
	ModbusVersion   int
	RatedPower      int
	AcOutputType    int
	SerialNumber    string
	ModelName       string
	Dsp1SwVersion   int
	Dsp2SwVersion   int
	DspSvnVersion   int
	ArmSwVersion    int
	ArmSvnVersion   int
	SoftwareVersion string
	ArmVersion      string
}

func (v ETInfoValues) AsJson(e *ETInfo) map[string]interface{} {
	return map[string]interface{}{
		e.ModbusVersion.Id:   v.ModbusVersion,
		e.RatedPower.Id:      v.RatedPower,
		e.AcOutputType.Id:    v.AcOutputType,
		e.SerialNumber.Id:    v.SerialNumber,
		e.ModelName.Id:       v.ModelName,
		e.Dsp1SwVersion.Id:   v.Dsp1SwVersion,
		e.Dsp2SwVersion.Id:   v.Dsp2SwVersion,
		e.DspSvnVersion.Id:   v.DspSvnVersion,
		e.ArmSwVersion.Id:    v.ArmSwVersion,
		e.ArmSvnVersion.Id:   v.ArmSvnVersion,
		e.SoftwareVersion.Id: v.SoftwareVersion,
		e.ArmVersion.Id:      v.ArmVersion,
	}
}

func (v ETInfoValues) SupportsEcoMode2() bool {
	if v.Dsp1SwVersion < 8 {
		return false
	}
	if v.Dsp2SwVersion < 8 {
		return false
	}
	if v.ArmSwVersion < 19 {
		return false
	}

	return true
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func CreateETInfo() *ETInfo {
	return &ETInfo{
		ModbusVersion: sensor.Integer("modbus_version", 0, "Modbus Version", "", sensor.NA),
		RatedPower:    sensor.Integer("rated_power", 2, "Rated Power", "", sensor.NA),

		// 0: 1-phase, 1: 3-phase (4 wire), 2: 3-phase (3 wire)
		AcOutputType: sensor.Integer("ac_output_type", 4, "AC Output Type", "", sensor.NA),

		SerialNumber: sensor.String("serial_number", "Serial Number", 6, 22),
		ModelName:    sensor.String("serial_number", "Serial Number", 22, 32),

		Dsp1SwVersion: sensor.Integer("dsp1_sw_version", 32, "Rated Power", "", sensor.NA),
		Dsp2SwVersion: sensor.Integer("dsp2_sw_version", 34, "Rated Power", "", sensor.NA),
		DspSvnVersion: sensor.Integer("dsp_svn_version", 36, "Rated Power", "", sensor.NA),
		ArmSwVersion:  sensor.Integer("arm_sw_version", 38, "Rated Power", "", sensor.NA),
		ArmSvnVersion: sensor.Integer("arm_svn_version", 40, "Rated Power", "", sensor.NA),

		SoftwareVersion: sensor.String("software_version", "Software Version", 42, 54),
		ArmVersion:      sensor.String("arm_version", "Arm Version", 54, 66),
	}
}
