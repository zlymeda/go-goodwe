package goodwe

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/zlymeda/go-goodwe/protocol"
)

type InverterResponse struct {
	Ip   net.IP
	Mac  string
	Wifi string
}

func SearchInverters(ctx context.Context) (InverterResponse, error) {
	t, err := NewTransport()
	if err != nil {
		return InverterResponse{}, err
	}

	return SearchInvertersWithTransport(ctx, t)
}

func SearchInvertersWithTransport(ctx context.Context, t *Transport) (InverterResponse, error) {
	broadcastAddress := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 48899,
	}

	req := protocol.NewSearchRequest()
	data, err := t.Execute(ctx, broadcastAddress, req)
	if err != nil {
		return InverterResponse{}, err
	}

	return parseSearchResponse(string(data))
}

func parseSearchResponse(resp string) (InverterResponse, error) {
	parts := strings.Split(resp, ",")
	if len(parts) != 3 {
		return InverterResponse{},
			fmt.Errorf("wrong inverter response: %s", resp)
	}

	ip, err := parseIp(parts[0])
	if err != nil {
		return InverterResponse{}, err
	}

	return InverterResponse{
		Ip:   ip,
		Mac:  parts[1],
		Wifi: parts[2],
	}, nil
}

func parseIp(ip string) (net.IP, error) {
	addr, err := net.ResolveIPAddr("ip", ip)

	if err != nil {
		return nil, err
	}

	return addr.IP, nil
}
