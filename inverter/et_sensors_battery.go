package inverter

import (
	"bytes"
	"github.com/zlymeda/go-goodwe"
	"github.com/zlymeda/go-goodwe/sensor"
)

//goland:noinspection SpellCheckingInspection
type ETSensorsBattery struct {
	Bms              sensor.Sensor[int16]
	Index            sensor.Sensor[int16]
	Status           sensor.Sensor[int16]
	Temperature      sensor.Sensor[float64]
	ChargeLimit      sensor.Sensor[int16]
	DischargeLimit   sensor.Sensor[int16]
	ErrorL           sensor.Sensor[int16]
	SOC              sensor.Sensor[int16]
	SOH              sensor.Sensor[int16]
	Modules          sensor.Sensor[int16]
	WarningL         sensor.Sensor[int16]
	Protocol         sensor.Sensor[int16]
	ErrorH           sensor.Sensor[int16]
	Error            sensor.Sensor[string]
	WarningH         sensor.Sensor[int16]
	Warning          sensor.Sensor[string]
	SwVersion        sensor.Sensor[int16]
	HwVersion        sensor.Sensor[int16]
	MaxCellTempId    sensor.Sensor[int16]
	MinCellTempId    sensor.Sensor[int16]
	MaxCellVoltageId sensor.Sensor[int16]
	MinCellVoltageId sensor.Sensor[int16]
	MaxCellTemp      sensor.Sensor[float64]
	MinCellTemp      sensor.Sensor[float64]
	MaxCellVoltage   sensor.Sensor[float64]
	MinCellVoltage   sensor.Sensor[float64]
}

func (s *ETSensorsBattery) Read(reader *bytes.Reader) ETSensorsBatteryValues {
	if reader == nil {
		return ETSensorsBatteryValues{}
	}

	return ETSensorsBatteryValues{
		Bms:              int(s.Bms.Read(reader)),
		Index:            int(s.Index.Read(reader)),
		Status:           int(s.Status.Read(reader)),
		Temperature:      s.Temperature.Read(reader),
		ChargeLimit:      int(s.ChargeLimit.Read(reader)),
		DischargeLimit:   int(s.DischargeLimit.Read(reader)),
		ErrorL:           int(s.ErrorL.Read(reader)),
		SOC:              int(s.SOC.Read(reader)),
		SOH:              int(s.SOH.Read(reader)),
		Modules:          int(s.Modules.Read(reader)),
		WarningL:         int(s.WarningL.Read(reader)),
		Protocol:         int(s.Protocol.Read(reader)),
		ErrorH:           int(s.ErrorH.Read(reader)),
		Error:            s.Error.Read(reader),
		WarningH:         int(s.WarningH.Read(reader)),
		Warning:          s.Warning.Read(reader),
		SwVersion:        int(s.SwVersion.Read(reader)),
		HwVersion:        int(s.HwVersion.Read(reader)),
		MaxCellTempId:    int(s.MaxCellTempId.Read(reader)),
		MinCellTempId:    int(s.MinCellTempId.Read(reader)),
		MaxCellVoltageId: int(s.MaxCellVoltageId.Read(reader)),
		MinCellVoltageId: int(s.MinCellVoltageId.Read(reader)),
		MaxCellTemp:      s.MaxCellTemp.Read(reader),
		MinCellTemp:      s.MinCellTemp.Read(reader),
		MaxCellVoltage:   s.MaxCellVoltage.Read(reader),
		MinCellVoltage:   s.MinCellVoltage.Read(reader),
	}
}

//goland:noinspection SpellCheckingInspection
type ETSensorsBatteryValues struct {
	Bms              int
	Index            int
	Status           int
	Temperature      float64
	ChargeLimit      int
	DischargeLimit   int
	ErrorL           int
	SOC              int
	SOH              int
	Modules          int
	WarningL         int
	Protocol         int
	ErrorH           int
	Error            string
	WarningH         int
	Warning          string
	SwVersion        int
	HwVersion        int
	MaxCellTempId    int
	MinCellTempId    int
	MaxCellVoltageId int
	MinCellVoltageId int
	MaxCellTemp      float64
	MinCellTemp      float64
	MaxCellVoltage   float64
	MinCellVoltage   float64
}

func (v ETSensorsBatteryValues) AsJson(e *ETSensorsBattery) map[string]interface{} {
	return map[string]interface{}{
		e.Bms.Id:              v.Bms,
		e.Index.Id:            v.Index,
		e.Status.Id:           v.Status,
		e.Temperature.Id:      v.Temperature,
		e.ChargeLimit.Id:      v.ChargeLimit,
		e.DischargeLimit.Id:   v.DischargeLimit,
		e.ErrorL.Id:           v.ErrorL,
		e.SOC.Id:              v.SOC,
		e.SOH.Id:              v.SOH,
		e.Modules.Id:          v.Modules,
		e.WarningL.Id:         v.WarningL,
		e.Protocol.Id:         v.Protocol,
		e.ErrorH.Id:           v.ErrorH,
		e.Error.Id:            v.Error,
		e.WarningH.Id:         v.WarningH,
		e.Warning.Id:          v.Warning,
		e.SwVersion.Id:        v.SwVersion,
		e.HwVersion.Id:        v.HwVersion,
		e.MaxCellTempId.Id:    v.MaxCellTempId,
		e.MinCellTempId.Id:    v.MinCellTempId,
		e.MaxCellVoltageId.Id: v.MaxCellVoltageId,
		e.MinCellVoltageId.Id: v.MinCellVoltageId,
		e.MaxCellTemp.Id:      v.MaxCellTemp,
		e.MinCellTemp.Id:      v.MinCellTemp,
		e.MaxCellVoltage.Id:   v.MaxCellVoltage,
		e.MinCellVoltage.Id:   v.MinCellVoltage,
	}
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func CreateETSensorsBattery() *ETSensorsBattery {

	errorH := sensor.Integer("battery_error_h", 24, "Battery Error H", "", sensor.BATTERY)
	errorL := sensor.Integer("battery_error_l", 12, "Battery Error L", "", sensor.BATTERY)

	warningH := sensor.Integer("battery_warning_h", 28, "Battery Warning H", "", sensor.BATTERY)
	warningL := sensor.Integer("battery_warning_l", 20, "Battery Warning L", "", sensor.BATTERY)

	return &ETSensorsBattery{
		Bms:            sensor.Integer("battery_bms", 0, "Battery BMS", "", sensor.BATTERY),
		Index:          sensor.Integer("battery_index", 2, "Battery Index", "", sensor.BATTERY),
		Status:         sensor.Integer("battery_status", 4, "Battery Status", "", sensor.BATTERY),
		Temperature:    sensor.Temperature("battery_temperature", 6, "Battery Temperature", sensor.BATTERY),
		ChargeLimit:    sensor.Integer("battery_charge_limit", 8, "Battery Charge Limit", "A", sensor.BATTERY),
		DischargeLimit: sensor.Integer("battery_discharge_limit", 10, "Battery Discharge Limit", "A", sensor.BATTERY),
		ErrorL:         errorL,
		SOC:            sensor.Integer("battery_soc", 14, "Battery State of Charge", "%", sensor.BATTERY),
		SOH:            sensor.Integer("battery_soh", 16, "Battery State of Health", "%", sensor.BATTERY),
		Modules:        sensor.Integer("battery_modules", 18, "Battery Modules", "", sensor.BATTERY),
		WarningL:       warningL,
		Protocol:       sensor.Integer("battery_protocol", 22, "Battery Protocol", "", sensor.BATTERY),
		ErrorH:         errorH,
		Error: sensor.Calculated("battery_error", "Battery Error", func(reader *bytes.Reader) string {
			return decodeBitmap(uint32(errorH.Read(reader))<<16+uint32(errorL.Read(reader)), goodwe.BmsAlarmCodes)
		}, "", sensor.BATTERY),
		WarningH: warningH,
		Warning: sensor.Calculated("battery_warning", "Battery Warning", func(reader *bytes.Reader) string {
			return decodeBitmap(uint32(warningH.Read(reader))<<16+uint32(warningL.Read(reader)), goodwe.BmsWarningCodes)
		}, "", sensor.BATTERY),
		SwVersion:        sensor.Integer("battery_sw_version", 30, "Battery Software Version", "", sensor.BATTERY),
		HwVersion:        sensor.Integer("battery_hw_version", 32, "Battery Hardware Version", "", sensor.BATTERY),
		MaxCellTempId:    sensor.Integer("battery_max_cell_temp_id", 34, "Battery Max Cell Temperature ID", "", sensor.BATTERY),
		MinCellTempId:    sensor.Integer("battery_min_cell_temp_id", 36, "Battery Min Cell Temperature ID", "", sensor.BATTERY),
		MaxCellVoltageId: sensor.Integer("battery_max_cell_voltage_id", 38, "Battery Max Cell Voltage ID", "", sensor.BATTERY),
		MinCellVoltageId: sensor.Integer("battery_min_cell_voltage_id", 40, "Battery Min Cell Voltage ID", "", sensor.BATTERY),
		MaxCellTemp:      sensor.Temperature("battery_max_cell_temp", 42, "Battery Max Cell Temperature", sensor.BATTERY),
		MinCellTemp:      sensor.Temperature("battery_min_cell_temp", 44, "Battery Min Cell Temperature", sensor.BATTERY),
		MaxCellVoltage:   sensor.Voltage("battery_max_cell_voltage", 46, "Battery Max Cell Voltage", sensor.BATTERY),
		MinCellVoltage:   sensor.Voltage("battery_min_cell_voltage", 48, "Battery Min Cell Voltage", sensor.BATTERY),
	}
}
