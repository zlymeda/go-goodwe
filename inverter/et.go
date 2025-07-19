package inverter

import (
	"bytes"
	"context"
	"github.com/zlymeda/go-goodwe"
	"github.com/zlymeda/go-goodwe/protocol"
	"github.com/zlymeda/go-goodwe/sensor"
	"log/slog"
	"net"
	"sync"
)

type ET struct {
	transport     *goodwe.Transport
	transportLock sync.Mutex

	inverter *net.UDPAddr

	comAddress uint8

	readDeviceInfo  protocol.Request
	readRunningData protocol.Request
	readMeterData   protocol.Request
	readBatteryInfo protocol.Request

	Sensors        *ETSensors
	SensorsBattery *ETSensorsBattery
	SensorsMeter   *ETSensorsMeter
	info           *ETInfo

	Settings *ETSettings
}

func NewETInverter(t *goodwe.Transport, inverter *net.UDPAddr) *ET {
	comAddress := uint8(0xf7)

	return &ET{
		transport:       t,
		inverter:        inverter,
		comAddress:      comAddress,
		readDeviceInfo:  protocol.NewModbusRead(comAddress, 0x88b8, 0x0021),
		readRunningData: protocol.NewModbusRead(comAddress, 0x891c, 0x007d),
		readMeterData:   protocol.NewModbusRead(comAddress, 0x8ca0, 0x2d),
		readBatteryInfo: protocol.NewModbusRead(comAddress, 0x9088, 0x0018),

		Sensors:        CreateETSensors(),
		SensorsBattery: CreateETSensorsBattery(),
		SensorsMeter:   CreateETSensorsMeter(),
		info:           CreateETInfo(),
		Settings:       CreateETSettings(),
	}
}

func (e *ET) ReadDeviceInfo(ctx context.Context) ETInfoValues {
	return e.info.Read(e.readFromInverter(ctx, e.readDeviceInfo))
}

func (e *ET) ReadRuntimeData(ctx context.Context) ETSensorValues {
	return e.Sensors.Read(e.readFromInverter(ctx, e.readRunningData))
}

func (e *ET) ReadBatteryData(ctx context.Context) ETSensorsBatteryValues {
	return e.SensorsBattery.Read(e.readFromInverter(ctx, e.readBatteryInfo))
}

func (e *ET) ReadMeterData(ctx context.Context) ETSensorsMeterValues {
	return e.SensorsMeter.Read(e.readFromInverter(ctx, e.readMeterData))
}

func (e *ET) ReadSettings(ctx context.Context) ETSettingsValues {
	return e.Settings.Read(e.ReadSettingBytes(ctx))
}

func (e *ET) ReadSettingBytes(ctx context.Context) func(s sensor.Sizeable) *bytes.Reader {
	return func(s sensor.Sizeable) *bytes.Reader {
		offset := uint16(s.GetOffset())
		size := uint16(s.GetSize())

		count := (size + (size % 2)) / 2

		return e.readFromInverter(ctx, protocol.NewModbusRead(e.comAddress, offset, count))
	}
}

func ReadSetting[T any](ctx context.Context, e *ET, setting sensor.Sensor[T]) T {
	return setting.JustRead(e.ReadSettingBytes(ctx)(setting))
}

func WriteSetting[T any](ctx context.Context, e *ET, setting sensor.Sensor[T], value T) {
	request := createWriteSettingRequest(e.comAddress, setting, value)
	e.readFromInverter(ctx, request)
}

func (e *ET) readFromInverter(ctx context.Context, request protocol.Request) *bytes.Reader {
	e.transportLock.Lock()
	defer e.transportLock.Unlock()

	response, err := e.transport.Execute(ctx, e.inverter, request)
	if err != nil {
		slog.Warn("readFromInverter", slog.Any("err", err))
		return nil
	}

	return bytes.NewReader(response[5 : len(response)-2])
}

func (e *ET) RuntimeAsJson(data ETSensorValues) map[string]interface{} {
	return data.AsJson(e.Sensors)
}

func (e *ET) BatteryAsJson(data ETSensorsBatteryValues) map[string]interface{} {
	return data.AsJson(e.SensorsBattery)
}

func (e *ET) InfoAsJson(data ETInfoValues) map[string]interface{} {
	return data.AsJson(e.info)
}

func (e *ET) MeterAsJson(data ETSensorsMeterValues) map[string]interface{} {
	return data.AsJson(e.SensorsMeter)
}

func (e *ET) SettingsAsJson(data ETSettingsValues) map[string]interface{} {
	return data.AsJson(e.Settings)
}

func createWriteSettingRequest[T any](comAddress uint8, setting sensor.Sensor[T], value T) protocol.Request {
	valueBytes := setting.Decode(value)

	if len(valueBytes) <= 2 {
		integerValue := sensor.ReadIntegerValue(bytes.NewReader(valueBytes))
		return protocol.NewModbusWrite(comAddress, uint16(setting.Offset), uint16(integerValue))
	}
	return protocol.NewModbusWriteMulti(comAddress, uint16(setting.Offset), valueBytes)
}
