package sensor

type Kind string

const (
	// PV inverter photo-voltaic (e.g. dc voltage of pv panels)
	PV Kind = "PV"

	// AC inverter grid output (e.g. ac voltage of grid connected output)
	AC Kind = "AC"

	// UPS inverter ups/eps/backup output (e.g. ac voltage of backup/off-grid connected output)
	UPS Kind = "UPS"

	// BATTERY (e.g. dc voltage of connected battery pack)
	BATTERY Kind = "BATTERY"

	// GRID power grid/smart meter (e.g. active power exported to grid)
	GRID Kind = "GRID"

	NA Kind = "N/A"
)
