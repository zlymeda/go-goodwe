package main

import (
	"context"
	"github.com/zlymeda/go-goodwe"
	"github.com/zlymeda/go-goodwe/internal/contexts"
	"github.com/zlymeda/go-goodwe/inverter"
	"github.com/zlymeda/go-goodwe/pkg/app"
	"github.com/zlymeda/go-goodwe/pkg/app/consumer"
	"log/slog"
	"os"
	"time"
)

func setupLogger() {
	stdout := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
	})

	logger := slog.New(stdout)
	slog.SetDefault(logger)
}

func main() {
	setupLogger()

	ctx, cancel := contexts.BackgroundSignalAwareContext()
	defer cancel()

	err := run(ctx)
	cancel()

	if err != nil {
		slog.Error("run err", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("goodwe done")
	os.Exit(0)
}

func run(ctx context.Context) error {
	myIp, err := goodwe.ResolveOutboundIp()
	if err != nil {
		return err
	}
	slog.Info("my ip", slog.String("ip", myIp.String()))

	et, cancel, err := app.CreateInverter(ctx, myIp)
	if err != nil {
		return err
	}

	defer cancel()

	events := make(chan inverter.Event)
	defer close(events)

	slog.Info("creating consumers..")
	consumers := []app.Consumer{
		consumer.CreateLogConsumer(),
		// can add more consumers here
	}

	slog.Info("creating consumers.. DONE")

	var consumersCh []chan inverter.Event
	for _, c := range consumers {
		ch := make(chan inverter.Event)
		consumersCh = append(consumersCh, ch)
		go c(ch)
	}

	go app.ChannelDistributor(events, consumersCh)

	app.ReadInverterData(ctx, et, events, 1*time.Second)

	return nil
}
