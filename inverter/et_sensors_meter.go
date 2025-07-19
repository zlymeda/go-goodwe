package inverter

import (
	"bytes"
	"github.com/zlymeda/go-goodwe/sensor"
)

//goland:noinspection SpellCheckingInspection
type ETSensorsMeter struct {
	Commode                 sensor.Sensor[int16]
	Rssi                    sensor.Sensor[int16]
	ManufactureCode         sensor.Sensor[int16]
	MeterTestStatus         sensor.Sensor[int16]
	MeterCommStatus         sensor.Sensor[int16]
	ActivePower1            sensor.Sensor[int16]
	ActivePower2            sensor.Sensor[int16]
	ActivePower3            sensor.Sensor[int16]
	ActivePower1Buy         sensor.Sensor[int16]
	ActivePower2Buy         sensor.Sensor[int16]
	ActivePower3Buy         sensor.Sensor[int16]
	ActivePower1Sell        sensor.Sensor[int16]
	ActivePower2Sell        sensor.Sensor[int16]
	ActivePower3Sell        sensor.Sensor[int16]
	ActivePowerTotal        sensor.Sensor[int16]
	ReactivePowerTotal      sensor.Sensor[int16]
	MeterPowerFactor1       sensor.Sensor[float64]
	MeterPowerFactor2       sensor.Sensor[float64]
	MeterPowerFactor3       sensor.Sensor[float64]
	MeterPowerFactor        sensor.Sensor[float64]
	MeterFreq               sensor.Sensor[float64]
	MeterETotalExp          sensor.Sensor[float64]
	MeterETotalImp          sensor.Sensor[float64]
	MeterActivePower1       sensor.Sensor[int32]
	MeterActivePower2       sensor.Sensor[int32]
	MeterActivePower3       sensor.Sensor[int32]
	MeterActivePowerTotal   sensor.Sensor[int32]
	MeterReactivePower1     sensor.Sensor[int32]
	MeterReactivePower2     sensor.Sensor[int32]
	MeterReactivePower3     sensor.Sensor[int32]
	MeterReactivePowerTotal sensor.Sensor[int32]
	MeterApparentPower1     sensor.Sensor[int32]
	MeterApparentPower2     sensor.Sensor[int32]
	MeterApparentPower3     sensor.Sensor[int32]
	MeterApparentPowerTotal sensor.Sensor[int32]
	MeterType               sensor.Sensor[int16]
	MeterSwVersion          sensor.Sensor[int16]
}

func (s *ETSensorsMeter) Read(reader *bytes.Reader) ETSensorsMeterValues {
	if reader == nil {
		return ETSensorsMeterValues{}
	}

	return ETSensorsMeterValues{
		Commode:                 int(s.Commode.Read(reader)),
		Rssi:                    int(s.Rssi.Read(reader)),
		ManufactureCode:         int(s.ManufactureCode.Read(reader)),
		MeterTestStatus:         int(s.MeterTestStatus.Read(reader)),
		MeterCommStatus:         int(s.MeterCommStatus.Read(reader)),
		ActivePower1:            int(s.ActivePower1.Read(reader)),
		ActivePower2:            int(s.ActivePower2.Read(reader)),
		ActivePower3:            int(s.ActivePower3.Read(reader)),
		ActivePower1Buy:         int(s.ActivePower1Buy.Read(reader)),
		ActivePower2Buy:         int(s.ActivePower2Buy.Read(reader)),
		ActivePower3Buy:         int(s.ActivePower3Buy.Read(reader)),
		ActivePower1Sell:        int(s.ActivePower1Sell.Read(reader)),
		ActivePower2Sell:        int(s.ActivePower2Sell.Read(reader)),
		ActivePower3Sell:        int(s.ActivePower3Sell.Read(reader)),
		ActivePowerTotal:        int(s.ActivePowerTotal.Read(reader)),
		ReactivePowerTotal:      int(s.ReactivePowerTotal.Read(reader)),
		MeterPowerFactor1:       s.MeterPowerFactor1.Read(reader),
		MeterPowerFactor2:       s.MeterPowerFactor2.Read(reader),
		MeterPowerFactor3:       s.MeterPowerFactor3.Read(reader),
		MeterPowerFactor:        s.MeterPowerFactor.Read(reader),
		MeterFreq:               s.MeterFreq.Read(reader),
		MeterETotalExp:          s.MeterETotalExp.Read(reader),
		MeterETotalImp:          s.MeterETotalImp.Read(reader),
		MeterActivePower1:       int(s.MeterActivePower1.Read(reader)),
		MeterActivePower2:       int(s.MeterActivePower2.Read(reader)),
		MeterActivePower3:       int(s.MeterActivePower3.Read(reader)),
		MeterActivePowerTotal:   int(s.MeterActivePowerTotal.Read(reader)),
		MeterReactivePower1:     int(s.MeterReactivePower1.Read(reader)),
		MeterReactivePower2:     int(s.MeterReactivePower2.Read(reader)),
		MeterReactivePower3:     int(s.MeterReactivePower3.Read(reader)),
		MeterReactivePowerTotal: int(s.MeterReactivePowerTotal.Read(reader)),
		MeterApparentPower1:     int(s.MeterApparentPower1.Read(reader)),
		MeterApparentPower2:     int(s.MeterApparentPower2.Read(reader)),
		MeterApparentPower3:     int(s.MeterApparentPower3.Read(reader)),
		MeterApparentPowerTotal: int(s.MeterApparentPowerTotal.Read(reader)),
		MeterType:               int(s.MeterType.Read(reader)),
		MeterSwVersion:          int(s.MeterSwVersion.Read(reader)),
	}
}

//goland:noinspection SpellCheckingInspection
type ETSensorsMeterValues struct {
	Commode                 int
	Rssi                    int
	ManufactureCode         int
	MeterTestStatus         int
	MeterCommStatus         int
	ActivePower1            int
	ActivePower2            int
	ActivePower3            int
	ActivePower1Buy         int
	ActivePower2Buy         int
	ActivePower3Buy         int
	ActivePower1Sell        int
	ActivePower2Sell        int
	ActivePower3Sell        int
	ActivePowerTotal        int
	ReactivePowerTotal      int
	MeterPowerFactor1       float64
	MeterPowerFactor2       float64
	MeterPowerFactor3       float64
	MeterPowerFactor        float64
	MeterFreq               float64
	MeterETotalExp          float64
	MeterETotalImp          float64
	MeterActivePower1       int
	MeterActivePower2       int
	MeterActivePower3       int
	MeterActivePowerTotal   int
	MeterReactivePower1     int
	MeterReactivePower2     int
	MeterReactivePower3     int
	MeterReactivePowerTotal int
	MeterApparentPower1     int
	MeterApparentPower2     int
	MeterApparentPower3     int
	MeterApparentPowerTotal int
	MeterType               int
	MeterSwVersion          int
}

func (v ETSensorsMeterValues) AsJson(e *ETSensorsMeter) map[string]interface{} {
	return map[string]interface{}{
		e.Commode.Id:                 v.Commode,
		e.Rssi.Id:                    v.Rssi,
		e.ManufactureCode.Id:         v.ManufactureCode,
		e.MeterTestStatus.Id:         v.MeterTestStatus,
		e.MeterCommStatus.Id:         v.MeterCommStatus,
		e.ActivePower1.Id:            v.ActivePower1,
		e.ActivePower2.Id:            v.ActivePower2,
		e.ActivePower3.Id:            v.ActivePower3,
		e.ActivePower1Sell.Id:        v.ActivePower1Sell,
		e.ActivePower2Sell.Id:        v.ActivePower2Sell,
		e.ActivePower3Sell.Id:        v.ActivePower3Sell,
		e.ActivePower1Buy.Id:         v.ActivePower1Buy,
		e.ActivePower2Buy.Id:         v.ActivePower2Buy,
		e.ActivePower3Buy.Id:         v.ActivePower3Buy,
		e.ActivePowerTotal.Id:        v.ActivePowerTotal,
		e.ReactivePowerTotal.Id:      v.ReactivePowerTotal,
		e.MeterPowerFactor1.Id:       v.MeterPowerFactor1,
		e.MeterPowerFactor2.Id:       v.MeterPowerFactor2,
		e.MeterPowerFactor3.Id:       v.MeterPowerFactor3,
		e.MeterPowerFactor.Id:        v.MeterPowerFactor,
		e.MeterFreq.Id:               v.MeterFreq,
		e.MeterETotalExp.Id:          v.MeterETotalExp,
		e.MeterETotalImp.Id:          v.MeterETotalImp,
		e.MeterActivePower1.Id:       v.MeterActivePower1,
		e.MeterActivePower2.Id:       v.MeterActivePower2,
		e.MeterActivePower3.Id:       v.MeterActivePower3,
		e.MeterActivePowerTotal.Id:   v.MeterActivePowerTotal,
		e.MeterReactivePower1.Id:     v.MeterReactivePower1,
		e.MeterReactivePower2.Id:     v.MeterReactivePower2,
		e.MeterReactivePower3.Id:     v.MeterReactivePower3,
		e.MeterReactivePowerTotal.Id: v.MeterReactivePowerTotal,
		e.MeterApparentPower1.Id:     v.MeterApparentPower1,
		e.MeterApparentPower2.Id:     v.MeterApparentPower2,
		e.MeterApparentPower3.Id:     v.MeterApparentPower3,
		e.MeterApparentPowerTotal.Id: v.MeterApparentPowerTotal,
		e.MeterType.Id:               v.MeterType,
		e.MeterSwVersion.Id:          v.MeterSwVersion,
	}
}

//goland:noinspection GoUnusedExportedFunction,SpellCheckingInspection
func CreateETSensorsMeter() *ETSensorsMeter {
	activePower1 := sensor.Power2("active_power1", 10, "Active Power L1", sensor.GRID)
	activePower2 := sensor.Power2("active_power2", 12, "Active Power L2", sensor.GRID)
	activePower3 := sensor.Power2("active_power3", 14, "Active Power L3", sensor.GRID)
	return &ETSensorsMeter{
		Commode:         sensor.Integer("commode", 0, "Commode", "", sensor.GRID),
		Rssi:            sensor.Integer("rssi", 2, "RSSI", "", sensor.GRID),
		ManufactureCode: sensor.Integer("manufacture_code", 4, "Manufacture Code", "", sensor.GRID),
		MeterTestStatus: sensor.Integer("meter_test_status", 6, "Meter Test Status", "", sensor.GRID),          // 1: correct，2: reverse，3: incorrect，0: not checked
		MeterCommStatus: sensor.Integer("meter_comm_status", 8, "Meter Communication Status", "", sensor.GRID), // 1 OK, 0 NotOK
		ActivePower1:    activePower1,
		ActivePower2:    activePower2,
		ActivePower3:    activePower3,

		ActivePower1Buy: sensor.Calculated(activePower1.Id+"_buy", activePower1.Name+" (Buy)", func(reader *bytes.Reader) int16 {
			// + sell, - buy
			power := activePower1.Read(reader)
			if power > 0 {
				return 0
			}
			return -power

		}, activePower1.Unit, activePower1.Kind),

		ActivePower1Sell: sensor.Calculated(activePower1.Id+"_sell", activePower1.Name+" (Sell)", func(reader *bytes.Reader) int16 {
			// + sell, - buy
			power := activePower1.Read(reader)
			if power < 0 {
				return 0
			}
			return power

		}, activePower1.Unit, activePower1.Kind),

		ActivePower2Buy: sensor.Calculated(activePower2.Id+"_buy", activePower2.Name+" (Buy)", func(reader *bytes.Reader) int16 {
			// + sell, - buy
			power := activePower2.Read(reader)
			if power > 0 {
				return 0
			}
			return -power

		}, activePower2.Unit, activePower2.Kind),

		ActivePower2Sell: sensor.Calculated(activePower2.Id+"_sell", activePower2.Name+" (Sell)", func(reader *bytes.Reader) int16 {
			// + sell, - buy
			power := activePower2.Read(reader)
			if power < 0 {
				return 0
			}
			return power

		}, activePower2.Unit, activePower2.Kind),

		ActivePower3Buy: sensor.Calculated(activePower3.Id+"_buy", activePower3.Name+" (Buy)", func(reader *bytes.Reader) int16 {
			// + sell, - buy
			power := activePower3.Read(reader)
			if power > 0 {
				return 0
			}
			return -power

		}, activePower3.Unit, activePower3.Kind),

		ActivePower3Sell: sensor.Calculated(activePower3.Id+"_sell", activePower3.Name+" (Sell)", func(reader *bytes.Reader) int16 {
			// + sell, - buy
			power := activePower3.Read(reader)
			if power < 0 {
				return 0
			}
			return power

		}, activePower3.Unit, activePower3.Kind),

		ActivePowerTotal:        sensor.Integer("active_power_total", 16, "Active Power Total", "W", sensor.GRID),
		ReactivePowerTotal:      sensor.Integer("reactive_power_total", 18, "Reactive Power Total", "var", sensor.GRID),
		MeterPowerFactor1:       sensor.Decimal("meter_power_factor1", 20, 1000, "Meter Power Factor L1", "", sensor.GRID),
		MeterPowerFactor2:       sensor.Decimal("meter_power_factor2", 22, 1000, "Meter Power Factor L2", "", sensor.GRID),
		MeterPowerFactor3:       sensor.Decimal("meter_power_factor3", 24, 1000, "Meter Power Factor L3", "", sensor.GRID),
		MeterPowerFactor:        sensor.Decimal("meter_power_factor", 26, 1000, "Meter Power Factor", "", sensor.GRID),
		MeterFreq:               sensor.Frequency("meter_freq", 28, "Meter Frequency", sensor.GRID),
		MeterETotalExp:          sensor.Float("meter_e_total_exp", 30, 1000, "Meter Total Energy (export)", "kWh", sensor.GRID),
		MeterETotalImp:          sensor.Float("meter_e_total_imp", 34, 1000, "Meter Total Energy (import)", "kWh", sensor.GRID),
		MeterActivePower1:       sensor.Long("meter_active_power1", 38, "Meter Active Power L1", "W", sensor.GRID),
		MeterActivePower2:       sensor.Long("meter_active_power2", 42, "Meter Active Power L2", "W", sensor.GRID),
		MeterActivePower3:       sensor.Long("meter_active_power3", 46, "Meter Active Power L3", "W", sensor.GRID),
		MeterActivePowerTotal:   sensor.Long("meter_active_power_total", 50, "Meter Active Power Total", "W", sensor.GRID),
		MeterReactivePower1:     sensor.Long("meter_reactive_power1", 54, "Meter Reactive Power L1", "var", sensor.GRID),
		MeterReactivePower2:     sensor.Long("meter_reactive_power2", 58, "Meter Reactive Power L2", "var", sensor.GRID),
		MeterReactivePower3:     sensor.Long("meter_reactive_power3", 62, "Meter Reactive Power L2", "var", sensor.GRID),
		MeterReactivePowerTotal: sensor.Long("meter_reactive_power_total", 66, "Meter Reactive Power Total", "var", sensor.GRID),
		MeterApparentPower1:     sensor.Long("meter_apparent_power1", 70, "Meter Apparent Power L1", "VA", sensor.GRID),
		MeterApparentPower2:     sensor.Long("meter_apparent_power2", 74, "Meter Apparent Power L2", "VA", sensor.GRID),
		MeterApparentPower3:     sensor.Long("meter_apparent_power3", 78, "Meter Apparent Power L3", "VA", sensor.GRID),
		MeterApparentPowerTotal: sensor.Long("meter_apparent_power_total", 82, "Meter Apparent Power Total", "VA", sensor.GRID),
		MeterType:               sensor.Integer("meter_type", 86, "Meter Type", "", sensor.GRID),
		MeterSwVersion:          sensor.Integer("meter_sw_version", 88, "Meter Software Version", "", sensor.GRID),
	}
}
