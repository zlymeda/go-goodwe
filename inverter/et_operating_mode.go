package inverter

import (
	"context"
	"fmt"
	"github.com/zlymeda/go-goodwe/protocol"
	"github.com/zlymeda/go-goodwe/sensor"
	"log/slog"
)

const (
	GeneralMode      OperatingMode = 0
	OffGridMode      OperatingMode = 1
	BackupMode       OperatingMode = 2
	EcoMode          OperatingMode = 3
	PeakShavingMode  OperatingMode = 4
	EcoModeCharge    OperatingMode = 5
	EcoModeDischarge OperatingMode = 6
	UnknownMode      OperatingMode = -1

	GeneralModeName      string = "General Mode"
	OffGridModeName      string = "Off-Grid Mode"
	BackupModeName       string = "Backup Mode"
	EcoModeName          string = "Eco Mode"
	PeakShavingModeName  string = "Peak Shaving Mode"
	EcoModeChargeName    string = "Eco Mode: Charge"
	EcoModeDischargeName string = "Eco Mode: Discharge"
)

type OperatingMode int

func (o OperatingMode) String() string {
	switch o {
	case GeneralMode:
		return GeneralModeName
	case OffGridMode:
		return OffGridModeName
	case BackupMode:
		return BackupModeName
	case EcoMode:
		return EcoModeName
	case PeakShavingMode:
		return PeakShavingModeName
	case EcoModeCharge:
		return EcoModeChargeName
	case EcoModeDischarge:
		return EcoModeDischargeName
	}

	panic(fmt.Sprintf("Unknown mode: %d", int(o)))
}

func OperatingModeFromString(mode string) OperatingMode {
	switch mode {
	case GeneralModeName:
		return GeneralMode
	case OffGridModeName:
		return OffGridMode
	case BackupModeName:
		return BackupMode
	case EcoModeName:
		return EcoMode
	case PeakShavingModeName:
		return PeakShavingMode
	case EcoModeChargeName:
		return EcoModeCharge
	case EcoModeDischargeName:
		return EcoModeDischarge
	}

	return UnknownMode
}

func (e *ET) GetOperatingMode(ctx context.Context) OperatingMode {
	setting := e.GetOperatingModeRaw(ctx)

	if setting != EcoMode {
		return setting
	}

	ecoSetting := ReadSetting(ctx, e, e.Settings.EcoModeV2S1)

	if ecoSetting.IsChargeMode() {
		return EcoModeCharge
	}
	if ecoSetting.IsDischargeMode() {
		return EcoModeDischarge
	}

	return EcoMode
}

func (e *ET) GetOperatingModeRaw(ctx context.Context) OperatingMode {
	return OperatingMode(ReadSetting(ctx, e, e.Settings.WorkMode))
}

func (e *ET) SetOperatingMode(ctx context.Context, mode OperatingMode, power uint, maxCharge uint) {

	switch mode {
	case GeneralMode:
		e.SetGeneralMode(ctx)

	case OffGridMode:
		e.SetOffGridMode(ctx)

	case BackupMode:
		e.SetBackupMode(ctx)

	case EcoMode:
		e.SetEcoMode(ctx)

	case PeakShavingMode:
		e.SetPeakShavingMode(ctx)

	case EcoModeCharge:
		e.SetEcoChargeMode(ctx, power, maxCharge)

	case EcoModeDischarge:
		e.SetEcoDischargeMode(ctx, power)
	}

}

func (e *ET) SetOffGridMode(ctx context.Context) {
	WriteSetting(ctx, e, e.Settings.WorkMode, int16(OffGridMode))
	e.setOffline(ctx, true)
	WriteSetting(ctx, e, e.Settings.BackupSupply, 1)
	WriteSetting(ctx, e, e.Settings.ColdStart, 4)
}

func (e *ET) SetGeneralMode(ctx context.Context) {
	e.setOnlineMode(ctx, GeneralMode)
}

func (e *ET) SetBackupMode(ctx context.Context) {
	e.setOnlineMode(ctx, BackupMode)
}

func (e *ET) setOnlineMode(ctx context.Context, mode OperatingMode) {
	WriteSetting(ctx, e, e.Settings.WorkMode, int16(mode))
	e.setOffline(ctx, false)
	e.clearBatteryModeParam(ctx)
}

func (e *ET) SetEcoMode(ctx context.Context) {
	WriteSetting(ctx, e, e.Settings.WorkMode, int16(EcoMode))
	e.setOffline(ctx, false)
}

func (e *ET) SetPeakShavingMode(ctx context.Context) {
	WriteSetting(ctx, e, e.Settings.WorkMode, 4)
	e.setOffline(ctx, false)
}

func (e *ET) SetEcoChargeMode(ctx context.Context, power uint, maxCharge uint) {
	if power < 0 || power > 100 {
		slog.Warn("power can be in [0, 100]", slog.Any("power", power))
		power = 50
	}
	if maxCharge < 0 || maxCharge > 100 {
		slog.Warn("maxCharge can be in [0, 100]", slog.Any("maxCharge", maxCharge))
		maxCharge = 100
	}

	e.writeEcoModeWithSettings(ctx, sensor.CreateChargeV2(power, maxCharge))
}

func (e *ET) SetEcoDischargeMode(ctx context.Context, power uint) {
	if power < 0 || power > 100 {
		slog.Warn("power can be in [0, 100]", slog.Any("power", power))
		power = 50
	}
	e.writeEcoModeWithSettings(ctx, sensor.CreateDischargeV2(power))
}

func (e *ET) writeEcoModeWithSettings(ctx context.Context, value sensor.ModeV2) {
	WriteSetting(ctx, e, e.Settings.EcoModeV2S1, value)
	WriteSetting(ctx, e, e.Settings.EcoModeV2S2, sensor.CreateOffV2())
	WriteSetting(ctx, e, e.Settings.EcoModeV2S3, sensor.CreateOffV2())
	WriteSetting(ctx, e, e.Settings.EcoModeV2S4, sensor.CreateOffV2())

	e.SetEcoMode(ctx)
}

func (e *ET) clearBatteryModeParam(ctx context.Context) {
	e.readFromInverter(ctx, protocol.NewModbusWrite(e.comAddress, 0xb9ad, 1))
}

func (e *ET) setOffline(ctx context.Context, offline bool) {
	value := int32(0x00010000)

	if offline {
		value = 0x00070000
	}

	bytez := sensor.EncodeLongValue(value)

	e.readFromInverter(ctx, protocol.NewModbusWriteMulti(e.comAddress, 0xb997, bytez))
}
