package goodwe

import (
	"context"
	"net"
)

type Service struct {
	transport *Transport
}

func NewService() (*Service, error) {
	transport, err := NewTransport()
	if err != nil {
		return nil, err
	}
	return NewServiceWithTransport(transport), nil
}

func NewServiceWithTransport(t *Transport) *Service {
	return &Service{transport: t}
}

func (s *Service) SearchInverters(ctx context.Context) (InverterResponse, error) {
	return SearchInvertersWithTransport(ctx, s.transport)
}

func (s *Service) Discover(ctx context.Context, inverter net.IP) (DiscoverResponse, error) {
	return DiscoverWithTransport(ctx, inverter, s.transport)
}

func (s *Service) GetTransport() *Transport {
	return s.transport
}
