package consumer

import (
	"fmt"
	"github.com/zlymeda/go-goodwe/inverter"
	"github.com/zlymeda/go-goodwe/pkg/app"
)

func CreateLogConsumer() app.Consumer {
	return func(events <-chan inverter.Event) {
		i := 0

		for e := range events {
			r := e.Runtime
			b := e.Battery

			if i == 0 {
				fmt.Println("date                   ,    PPV,  House,  Grid       ,  Battery,  SOC %, Bat. Label, Temp")
			}

			fmt.Printf("%s: %6d, %6d, %5d, %5s, %8.0f, %6d, %10s, %4.1f\n",
				e.Time.Format("2006-01-02 15:04:05.000"),
				r.Ppv, r.HouseConsumption, r.ActivePower, getGridLabel(r.ActivePower), r.Pbattery1, b.SOC, r.BatteryModeLabel, r.Temperature)

			i = (i + 1) % 50
		}
	}
}

func getGridLabel(activePower int) string {
	if activePower < 0 {
		return "Buy"

	}
	return "Sell"
}
