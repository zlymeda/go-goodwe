package app

import (
	"context"
	"fmt"
	"github.com/zlymeda/go-goodwe"
	"github.com/zlymeda/go-goodwe/inverter"
	"log/slog"
	"net"
	"time"
)

func ReadInverterData(ctx context.Context, et *inverter.ET, events chan<- inverter.Event, sleep time.Duration) {
	for {
		e := inverter.Event{
			Runtime: et.ReadRuntimeData(ctx),
			Battery: et.ReadBatteryData(ctx),
			Meter:   et.ReadMeterData(ctx),
		}

		e.Time = time.Now().UTC()

		if ctx.Err() != nil {
			return
		}

		select {
		case events <- e:
		case <-ctx.Done():
			return
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(sleep):
		}
	}
}

func CreateInverter(ctx context.Context, myIp *net.UDPAddr) (*inverter.ET, func(), error) {
	transport := goodwe.NewTransportWithAddr(myIp)

	scanCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	service := goodwe.NewServiceWithTransport(transport)

	inverterIp, err := getInverterIp(scanCtx, service)
	if err != nil {
		return nil, func() {}, fmt.Errorf("failed to get inverter IP: %w", err)
	}
	slog.Info("Inverter IP", slog.Any("ip", inverterIp))

	inverterAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", inverterIp, goodwe.UdpPort))
	if err != nil {
		return nil, func() {}, fmt.Errorf("failed to resolve inverter address: %w", err)
	}

	if err := transport.Open(); err != nil {
		return nil, func() {}, fmt.Errorf("failed to open transport: %w", err)
	}

	et := inverter.NewETInverter(transport, inverterAddress)

	return et, func() {
		_ = transport.Close()
	}, nil
}

func getInverterIp(ctx context.Context, service *goodwe.Service) (string, error) {
	resp, err := service.SearchInverters(ctx)
	if err != nil {
		return "", err
	}

	return resp.Ip.String(), nil
}
