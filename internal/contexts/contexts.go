package contexts

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Option func(s *ctxOpts)

type ctxOpts struct {
	timeoutBeforeKill time.Duration
}

func WithTimeoutBeforeKill(timeout time.Duration) Option {
	return func(s *ctxOpts) {
		s.timeoutBeforeKill = timeout
	}
}

func BackgroundSignalAwareContext(opts ...Option) (context.Context, func()) {
	var Signal = make(chan os.Signal, 1)
	signal.Notify(Signal,
		syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())

	main := ctxOpts{
		timeoutBeforeKill: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(&main)
	}

	go func() {
		sig := <-Signal
		slog.Info("Signal received", slog.Int("signal", int(sig.(syscall.Signal))))
		cancel()

		go func() {
			<-time.After(main.timeoutBeforeKill)
			slog.Info("Killed after timeout")
			os.Exit(6)
		}()

		go func() {
			sig := <-Signal
			slog.Info("Received 2nd signal; killing the server", slog.Int("signal", int(sig.(syscall.Signal))))
			os.Exit(5)
		}()
	}()

	return ctx, cancel
}
