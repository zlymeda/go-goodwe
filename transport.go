package goodwe

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/zlymeda/go-goodwe/protocol"
)

type Transport struct {
	local      *net.UDPAddr
	connection *net.UDPConn
	buffer     []byte
}

func NewTransport() (*Transport, error) {
	addr, err := ResolveOutboundIp()
	if err != nil {
		return nil, err
	}
	return NewTransportWithAddr(addr), nil
}

func NewTransportWithAddr(local *net.UDPAddr) *Transport {
	return &Transport{
		local:  local,
		buffer: make([]byte, 4096),
	}
}

func (t *Transport) Open() error {
	connection, err := net.ListenUDP("udp", t.local)
	if err != nil {
		return err
	}
	t.connection = connection
	return nil
}

func (t *Transport) Close() error {
	err := t.connection.Close()
	t.connection = nil
	return err
}

func (t *Transport) Execute(ctx context.Context, remote *net.UDPAddr, request protocol.Request) ([]byte, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if t.connection == nil {
		if err := t.Open(); err != nil {
			return nil, err
		}
		defer t.Close()
	}

	sendSignal := make(chan struct{}, 1)

	go func() {
		for {
			_ = t.connection.SetWriteDeadline(time.Now().Add(1 * time.Second))
			n, err := t.connection.WriteToUDP(request.Data, remote)
			if err != nil {
				if !isNetConnClosedErr(err) {
					slog.Warn("WriteToUDP err", slog.Any("err", err))
				}
				return
			}

			if n != len(request.Data) {
				slog.Warn("WriteToUDP: wrong number of bytes written",
					slog.Int("wanted", len(request.Data)),
					slog.Int("written", n),
					slog.Any("err", err),
				)
				continue
			}

			select {
			case <-time.After(500 * time.Millisecond):
			case <-sendSignal:
			case <-ctx.Done():
				return
			}
		}
	}()

	var partialResponse []byte
	for {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		_ = t.connection.SetReadDeadline(time.Now().Add(10 * time.Second))
		length, err := t.connection.Read(t.buffer)

		if err != nil {
			if errors.Is(err, os.ErrDeadlineExceeded) {
				continue
			}

			return nil, err
		}

		response := t.buffer[:length]
		if partialResponse != nil {
			response = append(partialResponse, response...)
		}

		err = request.ValidateResponse(response)
		if err == nil {
			partialResponse = nil
			return response, nil
		}

		var partialErr *protocol.PartialResponseErr
		if errors.As(err, &partialErr) {
			partialResponse = append(partialResponse, response...)
		} else {
			partialResponse = nil
		}

		select {
		case <-time.After(1 * time.Second):
		case sendSignal <- struct{}{}:
		case <-ctx.Done():
		}
	}
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

func ResolveOutboundIp() (*net.UDPAddr, error) {
	ip, err := GetOutboundIP()
	if err != nil {
		return nil, err
	}
	ourAddress, err := net.ResolveUDPAddr("udp", ip.String()+":0")
	if err != nil {
		return nil, err
	}
	return ourAddress, nil
}

func isNetConnClosedErr(err error) bool {
	switch {
	case
		errors.Is(err, net.ErrClosed),
		errors.Is(err, io.EOF),
		errors.Is(err, syscall.EPIPE):
		return true
	default:
		return false
	}
}
