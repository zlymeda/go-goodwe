package inverter

import (
	"bytes"
	"github.com/zlymeda/go-goodwe"
	"github.com/zlymeda/go-goodwe/sensor"
	"time"
)

//goland:noinspection SpellCheckingInspection
type ETSettings struct {
	Time                           sensor.Sensor[time.Time]
	CommAddress                    sensor.Sensor[int16]
	SensitivityCheck               sensor.Sensor[int16]
	ColdStart                      sensor.Sensor[int16]
	ShadowScan                     sensor.Sensor[int16]
	BackupSupply                   sensor.Sensor[int16]
	UnbalancedOutput               sensor.Sensor[int16]
	BatteryCapacity                sensor.Sensor[int16]
	BatteryModules                 sensor.Sensor[int16]
	BatteryChargeVoltage           sensor.Sensor[float64]
	BatteryChargeCurrent           sensor.Sensor[float64]
	BatteryDischargeVoltage        sensor.Sensor[float64]
	BatteryDischargeCurrent        sensor.Sensor[float64]
	BatteryDischargeDepth          sensor.Sensor[int16]
	BatteryDischargeVoltageOffline sensor.Sensor[float64]
	BatteryDischargeDepthOffline   sensor.Sensor[int16]
	PowerFactor                    sensor.Sensor[float64]
	WorkMode                       sensor.Sensor[int16]
	WorkModeLabel                  sensor.Sensor[string]
	BatterySocProtection           sensor.Sensor[int16]
	GridExport                     sensor.Sensor[int16]
	GridExportLimit                sensor.Sensor[int16]
	BatteryProtocolCode            sensor.Sensor[int16]
	EcoModeV1S1                    sensor.Sensor[sensor.Mode]
	EcoModeV1S2                    sensor.Sensor[sensor.Mode]
	EcoModeV1S3                    sensor.Sensor[sensor.Mode]
	EcoModeV1S4                    sensor.Sensor[sensor.Mode]
	EcoModeV2S1                    sensor.Sensor[sensor.ModeV2]
	EcoModeV2S2                    sensor.Sensor[sensor.ModeV2]
	EcoModeV2S3                    sensor.Sensor[sensor.ModeV2]
	EcoModeV2S4                    sensor.Sensor[sensor.ModeV2]
	EcoModeV2S5                    sensor.Sensor[sensor.ModeV2]
	EcoModeV2S6                    sensor.Sensor[sensor.ModeV2]
	EcoModeV2S7                    sensor.Sensor[sensor.ModeV2]
	FastCharging                   sensor.Sensor[int16]
	FastChargingSoc                sensor.Sensor[int16]
	FastChargingPower              sensor.Sensor[int16]
	DepthOfDischargeHolding        sensor.Sensor[int16]
	LoadControlMode                sensor.Sensor[int16]
	LoadControlSwitch              sensor.Sensor[int16]
	LoadControlSoC                 sensor.Sensor[int16]
}

func (s *ETSettings) Read(read func(s sensor.Sizeable) *bytes.Reader) ETSettingsValues {
	return ETSettingsValues{
		Time:                           s.Time.ReadUsing(read),
		CommAddress:                    int(s.CommAddress.ReadUsing(read)),
		SensitivityCheck:               int(s.SensitivityCheck.ReadUsing(read)),
		ColdStart:                      int(s.ColdStart.ReadUsing(read)),
		ShadowScan:                     int(s.ShadowScan.ReadUsing(read)),
		BackupSupply:                   int(s.BackupSupply.ReadUsing(read)),
		UnbalancedOutput:               int(s.UnbalancedOutput.ReadUsing(read)),
		BatteryCapacity:                int(s.BatteryCapacity.ReadUsing(read)),
		BatteryModules:                 int(s.BatteryModules.ReadUsing(read)),
		BatteryChargeVoltage:           s.BatteryChargeVoltage.ReadUsing(read),
		BatteryChargeCurrent:           s.BatteryChargeCurrent.ReadUsing(read),
		BatteryDischargeVoltage:        s.BatteryDischargeVoltage.ReadUsing(read),
		BatteryDischargeCurrent:        s.BatteryDischargeCurrent.ReadUsing(read),
		BatteryDischargeDepth:          int(s.BatteryDischargeDepth.ReadUsing(read)),
		BatteryDischargeVoltageOffline: s.BatteryDischargeVoltageOffline.ReadUsing(read),
		BatteryDischargeDepthOffline:   int(s.BatteryDischargeDepthOffline.ReadUsing(read)),
		PowerFactor:                    s.PowerFactor.ReadUsing(read),
		WorkMode:                       int(s.WorkMode.ReadUsing(read)),
		WorkModeLabel:                  s.WorkModeLabel.ReadUsing(read),
		BatterySocProtection:           int(s.BatterySocProtection.ReadUsing(read)),
		GridExport:                     int(s.GridExport.ReadUsing(read)),
		GridExportLimit:                int(s.GridExportLimit.ReadUsing(read)),
		EcoModeV1S1:                    s.EcoModeV1S1.ReadUsing(read),
		EcoModeV1S2:                    s.EcoModeV1S2.ReadUsing(read),
		EcoModeV1S3:                    s.EcoModeV1S3.ReadUsing(read),
		EcoModeV1S4:                    s.EcoModeV1S4.ReadUsing(read),
		EcoModeV2S1:                    s.EcoModeV2S1.ReadUsing(read),
		EcoModeV2S2:                    s.EcoModeV2S2.ReadUsing(read),
		EcoModeV2S3:                    s.EcoModeV2S3.ReadUsing(read),
		EcoModeV2S4:                    s.EcoModeV2S4.ReadUsing(read),
	}
}

//goland:noinspection SpellCheckingInspection
type ETSettingsValues struct {
	Time                           time.Time
	CommAddress                    int
	SensitivityCheck               int
	ColdStart                      int
	ShadowScan                     int
	BackupSupply                   int
	UnbalancedOutput               int
	BatteryCapacity                int
	BatteryModules                 int
	BatteryChargeVoltage           float64
	BatteryChargeCurrent           float64
	BatteryDischargeVoltage        float64
	BatteryDischargeCurrent        float64
	BatteryDischargeDepth          int
	BatteryDischargeVoltageOffline float64
	BatteryDischargeDepthOffline   int
	PowerFactor                    float64
	WorkMode                       int
	WorkModeLabel                  string
	BatterySocProtection           int
	GridExport                     int
	GridExportLimit                int
	EcoModeV1S1                    sensor.Mode
	EcoModeV1S2                    sensor.Mode
	EcoModeV1S3                    sensor.Mode
	EcoModeV1S4                    sensor.Mode
	EcoModeV2S1                    sensor.ModeV2
	EcoModeV2S2                    sensor.ModeV2
	EcoModeV2S3                    sensor.ModeV2
	EcoModeV2S4                    sensor.ModeV2
}

func (v ETSettingsValues) AsJson(e *ETSettings) map[string]interface{} {
	return map[string]interface{}{
		e.Time.Id:                           v.Time,
		e.CommAddress.Id:                    v.CommAddress,
		e.SensitivityCheck.Id:               v.SensitivityCheck,
		e.ColdStart.Id:                      v.ColdStart,
		e.ShadowScan.Id:                     v.ShadowScan,
		e.BackupSupply.Id:                   v.BackupSupply,
		e.UnbalancedOutput.Id:               v.UnbalancedOutput,
		e.BatteryCapacity.Id:                v.BatteryCapacity,
		e.BatteryModules.Id:                 v.BatteryModules,
		e.BatteryChargeVoltage.Id:           v.BatteryChargeVoltage,
		e.BatteryChargeCurrent.Id:           v.BatteryChargeCurrent,
		e.BatteryDischargeVoltage.Id:        v.BatteryDischargeVoltage,
		e.BatteryDischargeCurrent.Id:        v.BatteryDischargeCurrent,
		e.BatteryDischargeDepth.Id:          v.BatteryDischargeDepth,
		e.BatteryDischargeVoltageOffline.Id: v.BatteryDischargeVoltageOffline,
		e.BatteryDischargeDepthOffline.Id:   v.BatteryDischargeDepthOffline,
		e.PowerFactor.Id:                    v.PowerFactor,
		e.WorkMode.Id:                       v.WorkMode,
		e.WorkModeLabel.Id:                  v.WorkModeLabel,
		e.BatterySocProtection.Id:           v.BatterySocProtection,
		e.GridExport.Id:                     v.GridExport,
		e.GridExportLimit.Id:                v.GridExportLimit,
		e.EcoModeV1S1.Id:                    v.EcoModeV1S1,
		e.EcoModeV1S2.Id:                    v.EcoModeV1S2,
		e.EcoModeV1S3.Id:                    v.EcoModeV1S3,
		e.EcoModeV1S4.Id:                    v.EcoModeV1S4,
		e.EcoModeV2S1.Id:                    v.EcoModeV2S1,
		e.EcoModeV2S2.Id:                    v.EcoModeV2S2,
		e.EcoModeV2S3.Id:                    v.EcoModeV2S3,
		e.EcoModeV2S4.Id:                    v.EcoModeV2S4,
	}
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func CreateETSettings() *ETSettings {
	return &ETSettings{
		Time:                           sensor.Timestamp("time", 45200, "Inverter time"),
		CommAddress:                    sensor.Integer("comm_address", 45127, "Communication Address", "", sensor.NA),
		SensitivityCheck:               sensor.Integer("sensitivity_check", 45246, "Sensitivity Check Mode", "", sensor.AC),
		ColdStart:                      sensor.Integer("cold_start", 45248, "Cold Start", "", sensor.AC),
		ShadowScan:                     sensor.Integer("shadow_scan", 45251, "Shadow Scan", "", sensor.PV),
		BackupSupply:                   sensor.Integer("backup_supply", 45252, "Backup Supply", "", sensor.UPS),
		UnbalancedOutput:               sensor.Integer("unbalanced_output", 45264, "Unbalanced Output", "", sensor.AC),
		BatteryCapacity:                sensor.Integer("battery_capacity", 45350, "Battery Capacity", "Ah", sensor.BATTERY),
		BatteryModules:                 sensor.Integer("battery_modules", 45351, "Battery Modules", "", sensor.BATTERY),
		BatteryChargeVoltage:           sensor.Voltage("battery_charge_voltage", 45352, "Battery Charge Voltage", sensor.BATTERY),
		BatteryChargeCurrent:           sensor.Current("battery_charge_current", 45353, "Battery Charge Current", sensor.BATTERY),
		BatteryDischargeVoltage:        sensor.Voltage("battery_discharge_voltage", 45354, "Battery Discharge Voltage", sensor.BATTERY),
		BatteryDischargeCurrent:        sensor.Current("battery_discharge_current", 45355, "Battery Discharge Current", sensor.BATTERY),
		BatteryDischargeDepth:          sensor.Integer("battery_discharge_depth", 45356, "Battery Discharge Depth", "%", sensor.BATTERY),
		BatteryDischargeVoltageOffline: sensor.Voltage("battery_discharge_voltage_offline", 45357, "Battery Discharge Voltage (off-line)", sensor.BATTERY),
		BatteryDischargeDepthOffline:   sensor.Integer("battery_discharge_depth_offline", 45358, "Battery Discharge Depth (off-line)", "%", sensor.BATTERY),
		PowerFactor:                    sensor.Decimal("power_factor", 45482, 100, "Power Factor", "", sensor.NA),
		WorkMode:                       sensor.Integer("work_mode", 47000, "Work Mode", "", sensor.AC),
		WorkModeLabel:                  sensor.Enum2("work_mode_label", 47000, "Work Mode", goodwe.WorkModeSetting, sensor.NA),
		BatterySocProtection:           sensor.Integer("battery_soc_protection", 47500, "Battery SoC Protection", "", sensor.BATTERY),
		GridExport:                     sensor.Integer("grid_export", 47509, "Grid Export Enabled", "", sensor.GRID),
		GridExportLimit:                sensor.Integer("grid_export_limit", 47510, "Grid Export Limit", "W", sensor.GRID),
		BatteryProtocolCode:            sensor.Integer("battery_protocol_code", 47514, "Battery Protocol Code", "W", sensor.BATTERY),
		EcoModeV1S1:                    sensor.EcoMode("eco_mode_1", 47515, "Eco Mode Power Group 1"),
		EcoModeV1S2:                    sensor.EcoMode("eco_mode_2", 47519, "Eco Mode Power Group 2"),
		EcoModeV1S3:                    sensor.EcoMode("eco_mode_3", 47523, "Eco Mode Power Group 3"),
		EcoModeV1S4:                    sensor.EcoMode("eco_mode_4", 47527, "Eco Mode Power Group 4"),
		EcoModeV2S1:                    sensor.EcoModeV2("eco_modeV2_1", 47547, "Eco Mode Version 2 Power Group 1"),
		EcoModeV2S2:                    sensor.EcoModeV2("eco_modeV2_2", 47553, "Eco Mode Version 2 Power Group 2"),
		EcoModeV2S3:                    sensor.EcoModeV2("eco_modeV2_3", 47559, "Eco Mode Version 2 Power Group 3"),
		EcoModeV2S4:                    sensor.EcoModeV2("eco_modeV2_4", 47565, "Eco Mode Version 2 Power Group 4"),
		EcoModeV2S5:                    sensor.EcoModeV2("eco_modeV2_5", 47565, "Eco Mode Version 2 Power Group 5"),
		EcoModeV2S6:                    sensor.EcoModeV2("eco_modeV2_6", 47565, "Eco Mode Version 2 Power Group 6"),
		EcoModeV2S7:                    sensor.EcoModeV2("eco_modeV2_7", 47565, "Eco Mode Version 2 Power Group 7"),
		FastCharging:                   sensor.Integer("fast_charging", 47545, "Fast Charging Enabled", "", sensor.BATTERY),
		FastChargingSoc:                sensor.Integer("fast_charging_soc", 47546, "Fast Charging SoC", "%", sensor.BATTERY),
		FastChargingPower:              sensor.Integer("fast_charging_power", 47603, "Fast Charging Power", "W", sensor.BATTERY),
		DepthOfDischargeHolding:        sensor.Integer("dod_holding", 47602, "DoD Holding", "", sensor.BATTERY),
		LoadControlMode:                sensor.Integer("load_holding_mode", 47595, "Load Control Mode", "", sensor.AC),
		LoadControlSwitch:              sensor.Integer("load_holding_switch", 47596, "Load Control Switch", "", sensor.AC),
		LoadControlSoC:                 sensor.Integer("load_holding_soc", 47597, "Load Control SoC", "", sensor.AC),
	}
}
