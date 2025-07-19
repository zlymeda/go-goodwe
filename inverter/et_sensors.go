package inverter

import (
	"bytes"
	"fmt"
	"github.com/zlymeda/go-goodwe"
	"github.com/zlymeda/go-goodwe/sensor"
	"strings"
	"time"
)

//goland:noinspection SpellCheckingInspection
type ETSensors struct {
	Timestamp sensor.Sensor[time.Time]

	Ppv1 sensor.Sensor[int32]
	Ppv2 sensor.Sensor[int32]
	Ppv  sensor.Sensor[int32]

	Vpv1 sensor.Sensor[float64]
	Vpv2 sensor.Sensor[float64]
	Ipv1 sensor.Sensor[float64]
	Ipv2 sensor.Sensor[float64]

	Pv2Mode      sensor.Sensor[int8]
	Pv2ModeLabel sensor.Sensor[string]

	Pv1Mode      sensor.Sensor[int8]
	Pv1ModeLabel sensor.Sensor[string]

	Vgrid sensor.Sensor[float64]
	Igrid sensor.Sensor[float64]
	Fgrid sensor.Sensor[float64]
	Pgrid sensor.Sensor[int32]

	Vgrid2 sensor.Sensor[float64]
	Igrid2 sensor.Sensor[float64]
	Fgrid2 sensor.Sensor[float64]
	Pgrid2 sensor.Sensor[int32]

	Vgrid3 sensor.Sensor[float64]
	Igrid3 sensor.Sensor[float64]
	Fgrid3 sensor.Sensor[float64]
	Pgrid3 sensor.Sensor[int32]

	GridMode      sensor.Sensor[int16]
	GridModeLabel sensor.Sensor[string]

	TotalInverterPower sensor.Sensor[int32]

	ActivePower     sensor.Sensor[int32]
	ActivePowerBuy  sensor.Sensor[int32]
	ActivePowerSell sensor.Sensor[int32]
	GridInOut       sensor.Sensor[uint8]
	GridInOutLabel  sensor.Sensor[string]

	ReactivePower sensor.Sensor[int32]
	ApparentPower sensor.Sensor[int32]
	BackupV1      sensor.Sensor[float64]
	BackupI1      sensor.Sensor[float64]
	BackupF1      sensor.Sensor[float64]
	LoadMode1     sensor.Sensor[int16]
	BackupP1      sensor.Sensor[int32]
	BackupV2      sensor.Sensor[float64]
	BackupI2      sensor.Sensor[float64]
	BackupF2      sensor.Sensor[float64]
	LoadMode2     sensor.Sensor[int16]
	BackupP2      sensor.Sensor[int32]
	BackupV3      sensor.Sensor[float64]
	BackupI3      sensor.Sensor[float64]
	BackupF3      sensor.Sensor[float64]
	LoadMode3     sensor.Sensor[int16]
	BackupP3      sensor.Sensor[int32]

	LoadP1 sensor.Sensor[int32]
	LoadP2 sensor.Sensor[int32]
	LoadP3 sensor.Sensor[int32]

	LoadTotalP1 sensor.Sensor[int32]
	LoadTotalP2 sensor.Sensor[int32]
	LoadTotalP3 sensor.Sensor[int32]
	LoadTotal   sensor.Sensor[int32]

	Grid1     sensor.Sensor[int32]
	Grid2     sensor.Sensor[int32]
	Grid3     sensor.Sensor[int32]
	Grid1Sell sensor.Sensor[int32]
	Grid2Sell sensor.Sensor[int32]
	Grid3Sell sensor.Sensor[int32]
	Grid1Buy  sensor.Sensor[int32]
	Grid2Buy  sensor.Sensor[int32]
	Grid3Buy  sensor.Sensor[int32]

	BackupPTotal      sensor.Sensor[int32]
	LoadPTotal        sensor.Sensor[int32]
	UpsLoad           sensor.Sensor[int16]
	TemperatureAir    sensor.Sensor[float64]
	TemperatureModule sensor.Sensor[float64]
	Temperature       sensor.Sensor[float64]
	FunctionBit       sensor.Sensor[int16]
	BusVoltage        sensor.Sensor[float64]
	NbusVoltage       sensor.Sensor[float64]

	Vbattery1          sensor.Sensor[float64]
	Ibattery1          sensor.Sensor[float64]
	Pbattery1          sensor.Sensor[float64]
	Pbattery1Charge    sensor.Sensor[float64]
	Pbattery1Discharge sensor.Sensor[float64]

	BatteryMode        sensor.Sensor[int16]
	BatteryModeLabel   sensor.Sensor[string]
	WarningCode        sensor.Sensor[int16]
	SafetyCountry      sensor.Sensor[int16]
	SafetyCountryLabel sensor.Sensor[string]
	WorkMode           sensor.Sensor[int16]
	WorkModeLabel      sensor.Sensor[string]
	OperationMode      sensor.Sensor[int16]

	ErrorCodes sensor.Sensor[int32]
	Errors     sensor.Sensor[string]

	ETotal             sensor.Sensor[float64]
	EDay               sensor.Sensor[float64]
	ETotalExp          sensor.Sensor[float64]
	HTotal             sensor.Sensor[int32]
	EDayExp            sensor.Sensor[float64]
	ETotalImp          sensor.Sensor[float64]
	EDayImp            sensor.Sensor[float64]
	ELoadTotal         sensor.Sensor[float64]
	ELoadDay           sensor.Sensor[float64]
	EBatChargeTotal    sensor.Sensor[float64]
	EBatChargeDay      sensor.Sensor[float64]
	EBatDischargeTotal sensor.Sensor[float64]
	EBatDischargeDay   sensor.Sensor[float64]

	DiagnoseResult      sensor.Sensor[int32]
	DiagnoseResultLabel sensor.Sensor[string]

	HouseConsumption sensor.Sensor[int32]
}

func (s *ETSensors) Read(reader *bytes.Reader) ETSensorValues {
	if reader == nil {
		return ETSensorValues{}
	}

	return ETSensorValues{
		Timestamp:           s.Timestamp.Read(reader),
		Ppv1:                int(s.Ppv1.Read(reader)),
		Ppv2:                int(s.Ppv2.Read(reader)),
		Ppv:                 int(s.Ppv.Read(reader)),
		Vpv1:                s.Vpv1.Read(reader),
		Vpv2:                s.Vpv2.Read(reader),
		Ipv1:                s.Ipv1.Read(reader),
		Ipv2:                s.Ipv2.Read(reader),
		Pv2Mode:             int(s.Pv2Mode.Read(reader)),
		Pv2ModeLabel:        s.Pv2ModeLabel.Read(reader),
		Pv1Mode:             int(s.Pv1Mode.Read(reader)),
		Pv1ModeLabel:        s.Pv1ModeLabel.Read(reader),
		Vgrid:               s.Vgrid.Read(reader),
		Igrid:               s.Igrid.Read(reader),
		Fgrid:               s.Fgrid.Read(reader),
		Pgrid:               int(s.Pgrid.Read(reader)),
		Vgrid2:              s.Vgrid2.Read(reader),
		Igrid2:              s.Igrid2.Read(reader),
		Fgrid2:              s.Fgrid2.Read(reader),
		Pgrid2:              int(s.Pgrid2.Read(reader)),
		Vgrid3:              s.Vgrid3.Read(reader),
		Igrid3:              s.Igrid3.Read(reader),
		Fgrid3:              s.Fgrid3.Read(reader),
		Pgrid3:              int(s.Pgrid3.Read(reader)),
		GridMode:            int(s.GridMode.Read(reader)),
		GridModeLabel:       s.GridModeLabel.Read(reader),
		TotalInverterPower:  int(s.TotalInverterPower.Read(reader)),
		ActivePower:         int(s.ActivePower.Read(reader)),
		ActivePowerBuy:      int(s.ActivePowerBuy.Read(reader)),
		ActivePowerSell:     int(s.ActivePowerSell.Read(reader)),
		GridInOut:           uint(s.GridInOut.Read(reader)),
		GridInOutLabel:      s.GridInOutLabel.Read(reader),
		ReactivePower:       int(s.ReactivePower.Read(reader)),
		ApparentPower:       int(s.ApparentPower.Read(reader)),
		BackupV1:            s.BackupV1.Read(reader),
		BackupI1:            s.BackupI1.Read(reader),
		BackupF1:            s.BackupF1.Read(reader),
		LoadMode1:           int(s.LoadMode1.Read(reader)),
		BackupP1:            int(s.BackupP1.Read(reader)),
		BackupV2:            s.BackupV2.Read(reader),
		BackupI2:            s.BackupI2.Read(reader),
		BackupF2:            s.BackupF2.Read(reader),
		LoadMode2:           int(s.LoadMode2.Read(reader)),
		BackupP2:            int(s.BackupP2.Read(reader)),
		BackupV3:            s.BackupV3.Read(reader),
		BackupI3:            s.BackupI3.Read(reader),
		BackupF3:            s.BackupF3.Read(reader),
		LoadMode3:           int(s.LoadMode3.Read(reader)),
		BackupP3:            int(s.BackupP3.Read(reader)),
		LoadP1:              int(s.LoadP1.Read(reader)),
		LoadP2:              int(s.LoadP2.Read(reader)),
		LoadP3:              int(s.LoadP3.Read(reader)),
		LoadTotalP1:         int(s.LoadTotalP1.Read(reader)),
		LoadTotalP2:         int(s.LoadTotalP2.Read(reader)),
		LoadTotalP3:         int(s.LoadTotalP3.Read(reader)),
		Grid1:               int(s.Grid1.Read(reader)),
		Grid2:               int(s.Grid2.Read(reader)),
		Grid3:               int(s.Grid3.Read(reader)),
		Grid1Buy:            int(s.Grid1Buy.Read(reader)),
		Grid2Buy:            int(s.Grid2Buy.Read(reader)),
		Grid3Buy:            int(s.Grid3Buy.Read(reader)),
		Grid1Sell:           int(s.Grid1Sell.Read(reader)),
		Grid2Sell:           int(s.Grid2Sell.Read(reader)),
		Grid3Sell:           int(s.Grid3Sell.Read(reader)),
		LoadTotal:           int(s.LoadTotal.Read(reader)),
		BackupPTotal:        int(s.BackupPTotal.Read(reader)),
		LoadPTotal:          int(s.LoadPTotal.Read(reader)),
		UpsLoad:             int(s.UpsLoad.Read(reader)),
		TemperatureAir:      s.TemperatureAir.Read(reader),
		TemperatureModule:   s.TemperatureModule.Read(reader),
		Temperature:         s.Temperature.Read(reader),
		FunctionBit:         int(s.FunctionBit.Read(reader)),
		BusVoltage:          s.BusVoltage.Read(reader),
		NbusVoltage:         s.NbusVoltage.Read(reader),
		Vbattery1:           s.Vbattery1.Read(reader),
		Ibattery1:           s.Ibattery1.Read(reader),
		Pbattery1:           s.Pbattery1.Read(reader),
		Pbattery1Charge:     s.Pbattery1Charge.Read(reader),
		Pbattery1Discharge:  s.Pbattery1Discharge.Read(reader),
		BatteryMode:         int(s.BatteryMode.Read(reader)),
		BatteryModeLabel:    s.BatteryModeLabel.Read(reader),
		WarningCode:         int(s.WarningCode.Read(reader)),
		SafetyCountry:       int(s.SafetyCountry.Read(reader)),
		SafetyCountryLabel:  s.SafetyCountryLabel.Read(reader),
		WorkMode:            int(s.WorkMode.Read(reader)),
		WorkModeLabel:       s.WorkModeLabel.Read(reader),
		OperationMode:       int(s.OperationMode.Read(reader)),
		ErrorCodes:          int(s.ErrorCodes.Read(reader)),
		Errors:              s.Errors.Read(reader),
		ETotal:              s.ETotal.Read(reader),
		EDay:                s.EDay.Read(reader),
		ETotalExp:           s.ETotalExp.Read(reader),
		HTotal:              int(s.HTotal.Read(reader)),
		EDayExp:             s.EDayExp.Read(reader),
		ETotalImp:           s.ETotalImp.Read(reader),
		EDayImp:             s.EDayImp.Read(reader),
		ELoadTotal:          s.ELoadTotal.Read(reader),
		ELoadDay:            s.ELoadDay.Read(reader),
		EBatChargeTotal:     s.EBatChargeTotal.Read(reader),
		EBatChargeDay:       s.EBatChargeDay.Read(reader),
		EBatDischargeTotal:  s.EBatDischargeTotal.Read(reader),
		EBatDischargeDay:    s.EBatDischargeDay.Read(reader),
		DiagnoseResult:      int(s.DiagnoseResult.Read(reader)),
		DiagnoseResultLabel: s.DiagnoseResultLabel.Read(reader),
		HouseConsumption:    int(s.HouseConsumption.Read(reader)),
	}
}

//goland:noinspection SpellCheckingInspection
type ETSensorValues struct {
	Timestamp           time.Time
	Ppv1                int
	Ppv2                int
	Ppv                 int
	Vpv1                float64
	Vpv2                float64
	Ipv1                float64
	Ipv2                float64
	Pv2Mode             int
	Pv2ModeLabel        string
	Pv1Mode             int
	Pv1ModeLabel        string
	Vgrid               float64
	Igrid               float64
	Fgrid               float64
	Pgrid               int
	Vgrid2              float64
	Igrid2              float64
	Fgrid2              float64
	Pgrid2              int
	Vgrid3              float64
	Igrid3              float64
	Fgrid3              float64
	Pgrid3              int
	GridMode            int
	GridModeLabel       string
	TotalInverterPower  int
	ActivePower         int
	ActivePowerBuy      int
	ActivePowerSell     int
	GridInOut           uint
	GridInOutLabel      string
	ReactivePower       int
	ApparentPower       int
	BackupV1            float64
	BackupI1            float64
	BackupF1            float64
	LoadMode1           int
	BackupP1            int
	BackupV2            float64
	BackupI2            float64
	BackupF2            float64
	LoadMode2           int
	BackupP2            int
	BackupV3            float64
	BackupI3            float64
	BackupF3            float64
	LoadMode3           int
	BackupP3            int
	LoadP1              int
	LoadP2              int
	LoadP3              int
	LoadTotalP1         int
	LoadTotalP2         int
	LoadTotalP3         int
	LoadTotal           int
	Grid1               int
	Grid2               int
	Grid3               int
	Grid1Buy            int
	Grid2Buy            int
	Grid3Buy            int
	Grid1Sell           int
	Grid2Sell           int
	Grid3Sell           int
	BackupPTotal        int
	LoadPTotal          int
	UpsLoad             int
	TemperatureAir      float64
	TemperatureModule   float64
	Temperature         float64
	FunctionBit         int
	BusVoltage          float64
	NbusVoltage         float64
	Vbattery1           float64
	Ibattery1           float64
	Pbattery1           float64
	Pbattery1Charge     float64
	Pbattery1Discharge  float64
	BatteryMode         int
	BatteryModeLabel    string
	WarningCode         int
	SafetyCountry       int
	SafetyCountryLabel  string
	WorkMode            int
	WorkModeLabel       string
	OperationMode       int
	ErrorCodes          int
	Errors              string
	ETotal              float64
	EDay                float64
	ETotalExp           float64
	HTotal              int
	EDayExp             float64
	ETotalImp           float64
	EDayImp             float64
	ELoadTotal          float64
	ELoadDay            float64
	EBatChargeTotal     float64
	EBatChargeDay       float64
	EBatDischargeTotal  float64
	EBatDischargeDay    float64
	DiagnoseResult      int
	DiagnoseResultLabel string
	HouseConsumption    int
}

func (v ETSensorValues) AsJson(e *ETSensors) map[string]interface{} {
	return map[string]interface{}{
		e.Timestamp.Id:           v.Timestamp,
		e.Ppv1.Id:                v.Ppv1,
		e.Ppv2.Id:                v.Ppv2,
		e.Ppv.Id:                 v.Ppv,
		e.Vpv1.Id:                v.Vpv1,
		e.Vpv2.Id:                v.Vpv2,
		e.Ipv1.Id:                v.Ipv1,
		e.Ipv2.Id:                v.Ipv2,
		e.Pv2Mode.Id:             v.Pv2Mode,
		e.Pv2ModeLabel.Id:        v.Pv2ModeLabel,
		e.Pv1Mode.Id:             v.Pv1Mode,
		e.Pv1ModeLabel.Id:        v.Pv1ModeLabel,
		e.Vgrid.Id:               v.Vgrid,
		e.Igrid.Id:               v.Igrid,
		e.Fgrid.Id:               v.Fgrid,
		e.Pgrid.Id:               v.Pgrid,
		e.Vgrid2.Id:              v.Vgrid2,
		e.Igrid2.Id:              v.Igrid2,
		e.Fgrid2.Id:              v.Fgrid2,
		e.Pgrid2.Id:              v.Pgrid2,
		e.Vgrid3.Id:              v.Vgrid3,
		e.Igrid3.Id:              v.Igrid3,
		e.Fgrid3.Id:              v.Fgrid3,
		e.Pgrid3.Id:              v.Pgrid3,
		e.GridMode.Id:            v.GridMode,
		e.GridModeLabel.Id:       v.GridModeLabel,
		e.TotalInverterPower.Id:  v.TotalInverterPower,
		e.ActivePower.Id:         v.ActivePower,
		e.ActivePowerBuy.Id:      v.ActivePowerBuy,
		e.ActivePowerSell.Id:     v.ActivePowerSell,
		e.GridInOut.Id:           v.GridInOut,
		e.GridInOutLabel.Id:      v.GridInOutLabel,
		e.ReactivePower.Id:       v.ReactivePower,
		e.ApparentPower.Id:       v.ApparentPower,
		e.BackupV1.Id:            v.BackupV1,
		e.BackupI1.Id:            v.BackupI1,
		e.BackupF1.Id:            v.BackupF1,
		e.LoadMode1.Id:           v.LoadMode1,
		e.BackupP1.Id:            v.BackupP1,
		e.BackupV2.Id:            v.BackupV2,
		e.BackupI2.Id:            v.BackupI2,
		e.BackupF2.Id:            v.BackupF2,
		e.LoadMode2.Id:           v.LoadMode2,
		e.BackupP2.Id:            v.BackupP2,
		e.BackupV3.Id:            v.BackupV3,
		e.BackupI3.Id:            v.BackupI3,
		e.BackupF3.Id:            v.BackupF3,
		e.LoadMode3.Id:           v.LoadMode3,
		e.BackupP3.Id:            v.BackupP3,
		e.LoadP1.Id:              v.LoadP1,
		e.LoadP2.Id:              v.LoadP2,
		e.LoadP3.Id:              v.LoadP3,
		e.LoadTotalP1.Id:         v.LoadTotalP1,
		e.LoadTotalP2.Id:         v.LoadTotalP2,
		e.LoadTotalP3.Id:         v.LoadTotalP3,
		e.LoadTotal.Id:           v.LoadTotal,
		e.Grid1.Id:               v.Grid1,
		e.Grid2.Id:               v.Grid2,
		e.Grid3.Id:               v.Grid3,
		e.Grid1Buy.Id:            v.Grid1Buy,
		e.Grid2Buy.Id:            v.Grid2Buy,
		e.Grid3Buy.Id:            v.Grid3Buy,
		e.Grid1Sell.Id:           v.Grid1Sell,
		e.Grid2Sell.Id:           v.Grid2Sell,
		e.Grid3Sell.Id:           v.Grid3Sell,
		e.BackupPTotal.Id:        v.BackupPTotal,
		e.LoadPTotal.Id:          v.LoadPTotal,
		e.UpsLoad.Id:             v.UpsLoad,
		e.TemperatureAir.Id:      v.TemperatureAir,
		e.TemperatureModule.Id:   v.TemperatureModule,
		e.Temperature.Id:         v.Temperature,
		e.FunctionBit.Id:         v.FunctionBit,
		e.BusVoltage.Id:          v.BusVoltage,
		e.NbusVoltage.Id:         v.NbusVoltage,
		e.Vbattery1.Id:           v.Vbattery1,
		e.Ibattery1.Id:           v.Ibattery1,
		e.Pbattery1.Id:           v.Pbattery1,
		e.Pbattery1Charge.Id:     v.Pbattery1Charge,
		e.Pbattery1Discharge.Id:  v.Pbattery1Discharge,
		e.BatteryMode.Id:         v.BatteryMode,
		e.BatteryModeLabel.Id:    v.BatteryModeLabel,
		e.WarningCode.Id:         v.WarningCode,
		e.SafetyCountry.Id:       v.SafetyCountry,
		e.SafetyCountryLabel.Id:  v.SafetyCountryLabel,
		e.WorkMode.Id:            v.WorkMode,
		e.WorkModeLabel.Id:       v.WorkModeLabel,
		e.OperationMode.Id:       v.OperationMode,
		e.ErrorCodes.Id:          v.ErrorCodes,
		e.Errors.Id:              v.Errors,
		e.ETotal.Id:              v.ETotal,
		e.EDay.Id:                v.EDay,
		e.ETotalExp.Id:           v.ETotalExp,
		e.HTotal.Id:              v.HTotal,
		e.EDayExp.Id:             v.EDayExp,
		e.ETotalImp.Id:           v.ETotalImp,
		e.EDayImp.Id:             v.EDayImp,
		e.ELoadTotal.Id:          v.ELoadTotal,
		e.ELoadDay.Id:            v.ELoadDay,
		e.EBatChargeTotal.Id:     v.EBatChargeTotal,
		e.EBatChargeDay.Id:       v.EBatChargeDay,
		e.EBatDischargeTotal.Id:  v.EBatDischargeTotal,
		e.EBatDischargeDay.Id:    v.EBatDischargeDay,
		e.DiagnoseResult.Id:      v.DiagnoseResult,
		e.DiagnoseResultLabel.Id: v.DiagnoseResultLabel,
		e.HouseConsumption.Id:    v.HouseConsumption,
	}
}

//goland:noinspection SpellCheckingInspection
func CreateETSensors() *ETSensors {

	ppv1 := sensor.Power4("ppv1", 10, "PV1 Power", sensor.PV)
	ppv2 := sensor.Power4("ppv2", 18, "PV2 Power", sensor.PV)
	ppv := sensor.Plus("ppv", "PPV", ppv1, ppv2)

	activePower := sensor.Power4("active_power", 78, "Active Power", sensor.GRID)

	vbaterry1 := sensor.Voltage("vbattery1", 160, "Battery Voltage", sensor.BATTERY)
	ibattery1 := sensor.Current("ibattery1", 162, "Battery Current", sensor.BATTERY)
	pbattery := sensor.Mult("pbattery1", "Battery Power", "W", vbaterry1, ibattery1)

	errorCodes := sensor.Long("error_codes", 178, "Error Codes", "", sensor.NA)

	diagnoseResult := sensor.Long("diagnose_result", 240, "Diag Status Code", "", sensor.NA)

	loadP1 := sensor.Power4("load_p1", 126, "Load L1", sensor.AC)
	loadP2 := sensor.Power4("load_p2", 130, "Load L2", sensor.AC)
	loadP3 := sensor.Power4("load_p3", 134, "Load L3", sensor.AC)

	backupP1 := sensor.Power4("backup_p1", 98, "Back-up L1 Power", sensor.UPS)
	backupP2 := sensor.Power4("backup_p2", 110, "Back-up L2 Power", sensor.UPS)
	backupP3 := sensor.Power4("backup_p3", 122, "Back-up L3 Power", sensor.UPS)

	loadP1Total := sensor.Calculated("load_p1_total", "Load P1 Total", func(reader *bytes.Reader) int32 {
		return loadP1.Read(reader) + backupP1.Read(reader)
	}, loadP1.Unit, loadP1.Kind)

	loadP2Total := sensor.Calculated("load_p2_total", "Load P2 Total", func(reader *bytes.Reader) int32 {
		return loadP2.Read(reader) + backupP2.Read(reader)
	}, loadP2.Unit, loadP2.Kind)

	loadP3Total := sensor.Calculated("load_p3_total", "Load P3 Total", func(reader *bytes.Reader) int32 {
		return loadP3.Read(reader) + backupP3.Read(reader)
	}, loadP3.Unit, loadP3.Kind)

	pgrid1 := sensor.Power4("pgrid", 48, "On-grid L1 Power", sensor.AC)
	pgrid2 := sensor.Power4("pgrid2", 58, "On-grid L2 Power", sensor.AC)
	pgrid3 := sensor.Power4("pgrid3", 68, "On-grid L3 Power", sensor.AC)

	grid1 := sensor.Calculated("net_grid1", "Grid 1", func(reader *bytes.Reader) int32 {
		return pgrid1.Read(reader) - loadP1Total.Read(reader)
	}, loadP1.Unit, loadP1.Kind)
	grid2 := sensor.Calculated("net_grid2", "Grid 2", func(reader *bytes.Reader) int32 {
		return pgrid2.Read(reader) - loadP2Total.Read(reader)
	}, loadP1.Unit, loadP1.Kind)
	grid3 := sensor.Calculated("net_grid3", "Grid 3", func(reader *bytes.Reader) int32 {
		return pgrid3.Read(reader) - loadP3Total.Read(reader)
	}, loadP1.Unit, loadP1.Kind)
	return &ETSensors{
		Timestamp: sensor.Timestamp("timestamp", 0, "Timestamp"),

		Ppv1: ppv1,
		Ppv2: ppv2,
		Ppv:  ppv,

		Vpv1: sensor.Voltage("vpv1", 6, "PV1 Voltage", sensor.PV),
		Vpv2: sensor.Voltage("vpv2", 14, "PV2 Voltage", sensor.PV),

		Ipv1: sensor.Current("ipv1", 8, "PV1 Current", sensor.PV),
		Ipv2: sensor.Current("ipv2", 16, "PV2 Current", sensor.PV),

		Pv2Mode:      sensor.Byte("pv2_mode", 40, "PV2 Mode code", "", sensor.PV),
		Pv2ModeLabel: sensor.Enum("pv2_mode_label", 40, "PV2 Mode", goodwe.PvModes, sensor.PV),
		Pv1Mode:      sensor.Byte("pv1_mode", 41, "PV1 Mode code", "", sensor.PV),
		Pv1ModeLabel: sensor.Enum("pv1_mode_label", 41, "PV1 Mode", goodwe.PvModes, sensor.PV),

		Vgrid: sensor.Voltage("vgrid", 42, "On-grid L1 Voltage", sensor.AC),
		Igrid: sensor.Current("igrid", 44, "On-grid L1 Current", sensor.AC),
		Fgrid: sensor.Frequency("fgrid", 46, "On-grid L1 Frequency", sensor.AC),
		Pgrid: pgrid1,

		Vgrid2: sensor.Voltage("vgrid2", 52, "On-grid L2 Voltage", sensor.AC),
		Igrid2: sensor.Current("igrid2", 54, "On-grid L2 Current", sensor.AC),
		Fgrid2: sensor.Frequency("fgrid2", 56, "On-grid L2 Frequency", sensor.AC),
		Pgrid2: pgrid2,

		Vgrid3: sensor.Voltage("vgrid3", 62, "On-grid L3 Voltage", sensor.AC),
		Igrid3: sensor.Current("igrid3", 64, "On-grid L3 Current", sensor.AC),
		Fgrid3: sensor.Frequency("fgrid3", 66, "On-grid L3 Frequency", sensor.AC),
		Pgrid3: pgrid3,

		GridMode:      sensor.Integer("grid_mode", 72, "Grid Mode code", "", sensor.PV),
		GridModeLabel: sensor.Enum2("grid_mode", 72, "Grid Mode", goodwe.GridModes, sensor.PV),

		TotalInverterPower: sensor.Power4("total_inverter_power", 74, "Total Power", sensor.AC),
		ActivePower:        activePower,
		ActivePowerBuy: sensor.Calculated("active_power_buy", "Active Power (Buy)", func(reader *bytes.Reader) int32 {
			return buy(activePower.Read(reader))

		}, "W", sensor.GRID),
		ActivePowerSell: sensor.Calculated("active_power_sell", "Active Power (Sell)", func(reader *bytes.Reader) int32 {
			return sell(activePower.Read(reader))
		}, "W", sensor.GRID),

		GridInOut: sensor.Calculated("grid_in_out", "On-grid Mode code", func(reader *bytes.Reader) uint8 {
			return decodeGridMode(activePower.Read(reader))
		}, "", sensor.NA),

		GridInOutLabel: sensor.Calculated("grid_in_out", "On-grid Mode", func(reader *bytes.Reader) string {
			return goodwe.GridInOutModes[decodeGridMode(activePower.Read(reader))]
		}, "", sensor.NA),

		ReactivePower: sensor.Long("reactive_power", 82, "Reactive Power", "var", sensor.GRID),
		ApparentPower: sensor.Long("apparent_power", 86, "Apparent Power", "VA", sensor.GRID),
		BackupV1:      sensor.Voltage("backup_v1", 90, "Back-up L1 Voltage", sensor.UPS),
		BackupI1:      sensor.Current("backup_i1", 92, "Back-up L1 Current", sensor.UPS),
		BackupF1:      sensor.Frequency("backup_f1", 94, "Back-up L1 Frequency", sensor.UPS),
		LoadMode1:     sensor.Integer("load_mode1", 96, "Load Mode L1", "", sensor.NA),
		BackupP1:      backupP1,
		BackupV2:      sensor.Voltage("backup_v2", 102, "Back-up L2 Voltage", sensor.UPS),
		BackupI2:      sensor.Current("backup_i2", 104, "Back-up L2 Current", sensor.UPS),
		BackupF2:      sensor.Frequency("backup_f2", 106, "Back-up L2 Frequency", sensor.UPS),
		LoadMode2:     sensor.Integer("load_mode2", 108, "Load Mode L2", "", sensor.NA),
		BackupP2:      backupP2,
		BackupV3:      sensor.Voltage("backup_v3", 114, "Back-up L3 Voltage", sensor.UPS),
		BackupI3:      sensor.Current("backup_i3", 116, "Back-up L3 Current", sensor.UPS),
		BackupF3:      sensor.Frequency("backup_f3", 118, "Back-up L3 Frequency", sensor.UPS),
		LoadMode3:     sensor.Integer("load_mode3", 120, "Load Mode L3", "", sensor.NA),
		BackupP3:      backupP3,
		LoadP1:        loadP1,
		LoadP2:        loadP2,
		LoadP3:        loadP3,
		BackupPTotal:  sensor.Power4("backup_ptotal", 138, "Back-up Load", sensor.UPS),
		LoadPTotal:    sensor.Power4("load_ptotal", 142, "Load", sensor.AC),
		UpsLoad:       sensor.Integer("ups_load", 146, "Ups Load", "%", sensor.UPS),

		TemperatureAir:    sensor.Temperature("temperature_air", 148, "Inverter Temperature (Air)", sensor.AC),
		TemperatureModule: sensor.Temperature("temperature_module", 150, "Inverter Temperature (Module)", sensor.NA),
		Temperature:       sensor.Temperature("temperature", 152, "Inverter Temperature (Radiator)", sensor.AC),

		FunctionBit: sensor.Integer("function_bit", 154, "Function Bit", "", sensor.NA),
		BusVoltage:  sensor.Voltage("bus_voltage", 156, "Bus Voltage", sensor.NA),
		NbusVoltage: sensor.Voltage("nbus_voltage", 158, "NBus Voltage", sensor.NA),

		Vbattery1: vbaterry1,
		Ibattery1: ibattery1,
		Pbattery1: pbattery,
		Pbattery1Charge: sensor.Calculated("pbattery1_charge", "Battery Power (Charge)", func(reader *bytes.Reader) float64 {
			// Battery -C/+D
			power := pbattery.Read(reader)
			if power > 0 {
				return 0
			}
			return -power

		}, "W", sensor.BATTERY),
		Pbattery1Discharge: sensor.Calculated("pbattery1_discharge", "Battery Power (Discharge)", func(reader *bytes.Reader) float64 {
			// Battery -C/+D
			power := pbattery.Read(reader)
			if power > 0 {
				return power
			}
			return 0

		}, "W", sensor.BATTERY),

		BatteryMode:        sensor.Integer("battery_mode", 168, "Battery Mode code", "", sensor.BATTERY),
		BatteryModeLabel:   sensor.Enum2("battery_mode", 168, "Battery Mode", goodwe.BatteryModesEt, sensor.BATTERY),
		WarningCode:        sensor.Integer("warning_code", 170, "Warning code", "", sensor.NA),
		SafetyCountry:      sensor.Integer("safety_country", 172, "Safety Country code", "", sensor.AC),
		SafetyCountryLabel: sensor.Enum2("safety_country", 172, "Safety Country", goodwe.SafetyCountriesEt, sensor.AC),
		WorkMode:           sensor.Integer("work_mode", 174, "Work Mode code", "", sensor.AC),
		WorkModeLabel:      sensor.Enum2("work_mode", 174, "Work Mode", goodwe.WorkModesEt, sensor.NA),
		OperationMode:      sensor.Integer("operation_mode", 176, "Operation Mode code", "", sensor.NA),
		ErrorCodes:         errorCodes,
		Errors: sensor.Calculated(errorCodes.Id, "Errors", func(reader *bytes.Reader) string {
			return decodeBitmap(uint32(errorCodes.Read(reader)), goodwe.ErrorCodes)
		}, "", sensor.NA),

		ETotal:             sensor.Energy4("e_total", 182, "Total PV Generation", sensor.PV),
		EDay:               sensor.Energy4("e_day", 186, "Today's PV Generation", sensor.PV),
		ETotalExp:          sensor.Energy4("e_total_exp", 190, "Total Energy (export)", sensor.AC),
		HTotal:             sensor.Long("h_total", 194, "Hours Total", "h", sensor.PV),
		EDayExp:            sensor.Energy2("e_day_exp", 198, "Today Energy (export)", sensor.AC),
		ETotalImp:          sensor.Energy4("e_total_imp", 200, "Total Energy (import)", sensor.AC),
		EDayImp:            sensor.Energy2("e_day_imp", 204, "Today Energy (import)", sensor.AC),
		ELoadTotal:         sensor.Energy4("e_load_total", 206, "Total Load", sensor.AC),
		ELoadDay:           sensor.Energy2("e_load_day", 210, "Today Load", sensor.AC),
		EBatChargeTotal:    sensor.Energy4("e_bat_charge_total", 212, "Total Battery Charge", sensor.BATTERY),
		EBatChargeDay:      sensor.Energy2("e_bat_charge_day", 216, "Today Battery Charge", sensor.BATTERY),
		EBatDischargeTotal: sensor.Energy4("e_bat_discharge_total", 218, "Total Battery Discharge", sensor.BATTERY),
		EBatDischargeDay:   sensor.Energy2("e_bat_discharge_day", 222, "Today Battery Discharge", sensor.BATTERY),

		DiagnoseResult: diagnoseResult,
		DiagnoseResultLabel: sensor.Calculated(diagnoseResult.Id, "Errors", func(reader *bytes.Reader) string {
			return decodeBitmap(uint32(diagnoseResult.Read(reader)), goodwe.DiagStatusCodes)
		}, "", sensor.NA),

		HouseConsumption: sensor.Calculated("house_consumption", "House Consumption", func(reader *bytes.Reader) int32 {
			return int32(float64(ppv.Read(reader)) + pbattery.Read(reader) - float64(activePower.Read(reader)))
		}, "W", sensor.AC),

		LoadTotalP1: loadP1Total,
		LoadTotalP2: loadP2Total,
		LoadTotalP3: loadP3Total,

		Grid1: grid1,
		Grid2: grid2,
		Grid3: grid3,

		Grid1Sell: sensor.Calculated("net_grid1sell", "Grid 1 (Sell)", func(reader *bytes.Reader) int32 {
			return sell(grid1.Read(reader))
		}, loadP1.Unit, loadP1.Kind),
		Grid2Sell: sensor.Calculated("net_grid2sell", "Grid 2 (Sell)", func(reader *bytes.Reader) int32 {
			return sell(grid2.Read(reader))
		}, loadP1.Unit, loadP1.Kind),
		Grid3Sell: sensor.Calculated("net_grid3sell", "Grid 3 (Sell)", func(reader *bytes.Reader) int32 {
			return sell(grid3.Read(reader))
		}, loadP1.Unit, loadP1.Kind),

		Grid1Buy: sensor.Calculated("net_grid1buy", "Grid 1 (Buy)", func(reader *bytes.Reader) int32 {
			return buy(grid1.Read(reader))
		}, loadP1.Unit, loadP1.Kind),
		Grid2Buy: sensor.Calculated("net_grid2buy", "Grid 2 (Buy)", func(reader *bytes.Reader) int32 {
			return buy(grid2.Read(reader))
		}, loadP1.Unit, loadP1.Kind),
		Grid3Buy: sensor.Calculated("net_grid3buy", "Grid 3 (Buy)", func(reader *bytes.Reader) int32 {
			return buy(grid3.Read(reader))
		}, loadP1.Unit, loadP1.Kind),

		LoadTotal: sensor.Calculated("load_total", "Load Total", func(reader *bytes.Reader) int32 {
			return loadP1Total.Read(reader) + loadP2Total.Read(reader) + loadP3Total.Read(reader)
		}, loadP3.Unit, loadP3.Kind),
	}
}

func sell(power int32) int32 {
	// + sell, - buy
	if power < 0 {
		return 0
	}
	return power
}

func buy(power int32) int32 {
	// + sell, - buy
	if power > 0 {
		return 0
	}
	return -power
}

func decodeBitmap(value uint32, bitmap map[int]string) string {
	bits := value
	var result []string
	for i := 0; i < 32; i++ {
		if bits&0x1 == 1 {
			if val, ok := bitmap[i]; ok {
				result = append(result, val)
			} else {
				result = append(result, fmt.Sprintf("err%d", i))
			}
		}
		bits = bits >> 1
	}
	return strings.Join(result, ", ")
}

func decodeGridMode(value int32) uint8 {
	if value < -90 {
		return 2
	}
	if value >= 90 {
		return 1
	}

	return 0
}
