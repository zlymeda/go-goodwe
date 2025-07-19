package goodwe

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/zlymeda/go-goodwe/protocol"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	etFamily    = []string{"ET", "EH", "BT", "BH"}
	etModelTags = []string{"ETU", "EHU", "BTU", "BHU"}

	esFamily    = []string{"ES", "EM", "BP"}
	esModelTags = []string{"ESU", "EMU", "BPU", "BPS"}

	dtFamily    = []string{"DT", "MS", "NS", "XS"}
	dtModelTags = []string{"DTU", "MSU", "DTN", "DSN", "PSB", "PSC"}
)

type DiscoverResponse struct {
	ModelName    string
	SerialNumber string
}

func Discover(ctx context.Context, inverter net.IP) (DiscoverResponse, error) {
	t, err := NewTransport()
	if err != nil {
		return DiscoverResponse{}, err
	}

	return DiscoverWithTransport(ctx, inverter, t)
}

func DiscoverWithTransport(ctx context.Context, inverter net.IP, t *Transport) (DiscoverResponse, error) {
	inverterAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", inverter, UdpPort))
	if err != nil {
		return DiscoverResponse{}, err
	}

	req := protocol.NewDiscoverRequest()
	data, err := t.Execute(ctx, inverterAddress, req)
	if err != nil {
		return DiscoverResponse{}, err
	}

	return parseDiscoverResponse(data), nil
}

func parseDiscoverResponse(response []byte) DiscoverResponse {
	return DiscoverResponse{
		ModelName:    strings.TrimSpace(DecodeAscii(response[12:22])),
		SerialNumber: strings.TrimSpace(DecodeAscii(response[38:54])),
	}
}

func DecodeAscii(response []byte) string {
	r := ""
	for _, b := range response {
		r += fmt.Sprintf("%c", b)
	}
	return strings.TrimSpace(r)
}
