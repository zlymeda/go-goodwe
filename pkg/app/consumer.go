package app

import (
	"github.com/zlymeda/go-goodwe/inverter"
)

type Consumer func(<-chan inverter.Event)
