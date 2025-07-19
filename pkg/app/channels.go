package app

import "github.com/zlymeda/go-goodwe/inverter"

func ChannelDistributor(eventCh <-chan inverter.Event, consumers []chan inverter.Event) {
	for event := range eventCh {
		for _, ch := range consumers {
			// Non-blocking send to each consumer channel
			select {
			case ch <- event:
			default:
				// If the consumer channel is full, discard the event
			}
		}
	}

	for _, ch := range consumers {
		close(ch)
	}
}
