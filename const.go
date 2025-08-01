package goodwe

const UdpPort = 8899

var BatteryModesEt = map[uint16]string{
	0: "No battery",
	1: "Standby",
	2: "Discharge",
	3: "Charge",
	4: "To be charged",
	5: "To be discharged",
}

var EnergyModes = map[uint8]string{
	0:   "Check Mode",
	1:   "Wait Mode",
	2:   "Normal (On-Grid)",
	4:   "Normal (Off-Grid)",
	8:   "Flash Mode",
	16:  "Fault Mode",
	32:  "Battery Standby",
	64:  "Battery Charging",
	128: "Battery Discharging",
}

var GridModes = map[uint16]string{
	0: "Not connected to grid",
	1: "Connected to grid",
	2: "Fault",
}

var GridInOutModes = map[uint8]string{
	0: "Idle",
	1: "Exporting",
	2: "Importing",
}

var LoadModes = map[uint8]string{
	0: "Inverter and the load is disconnected",
	1: "The inverter is connected to a load",
}

var PvModes = map[uint8]string{
	0: "PV panels not connected",
	1: "PV panels connected, no power",
	2: "PV panels connected, producing power",
}

var WorkModes = map[uint8]string{
	0: "Wait Mode",
	1: "Normal",
	2: "Error",
	4: "Check Mode",
}

var WorkModesEt = map[uint16]string{
	0: "Wait Mode",
	1: "Normal (On-Grid)",
	2: "Normal (Off-Grid)",
	3: "Fault Mode",
	4: "Flash Mode",
	5: "Check Mode",
}
var WorkModesEs = map[uint8]string{
	0: "Inverter Off - Standby",
	1: "Inverter On",
	2: "Inverter Abnormal, stopping power",
	3: "Inverter Severely Abnormal, 20 seconds to restart",
}

const (
	GeneralMode     = 0
	OffGridMode     = 1
	BackupMode      = 2
	EcoMode         = 3
	PeakShavingMode = 4
)

var WorkModeSetting = map[uint16]string{
	GeneralMode:     "General mode",
	OffGridMode:     "Off grid mode",
	BackupMode:      "Backup mode",
	EcoMode:         "Eco mode",
	PeakShavingMode: "Peak shaving mode",
}

//goland:noinspection SpellCheckingInspection
var SafetyCountriesEt = map[uint16]string{
	0:   "Italy",
	1:   "Czechia",
	2:   "Germany",
	3:   "Spain",
	4:   "Greece",
	5:   "Denmark",
	6:   "Belgium",
	7:   "Romania",
	8:   "G98",
	9:   "Australia",
	10:  "France",
	11:  "China",
	13:  "Poland",
	14:  "South Africa",
	15:  "AustraliaL",
	16:  "Brazil",
	17:  "Thailand MEA",
	18:  "Thailand PEA",
	19:  "Mauritius",
	20:  "Holland",
	21:  "Northern Ireland",
	22:  "China Higher",
	23:  "French 50Hz",
	24:  "French 60Hz",
	25:  "Australia Ergon",
	26:  "Australia Energex",
	27:  "Holland 16/20A",
	28:  "Korea",
	29:  "China Station",
	30:  "Austria",
	31:  "India",
	32:  "50Hz Grid Default",
	33:  "Warehouse",
	34:  "Philippines",
	35:  "Ireland",
	36:  "Taiwan",
	37:  "Bulgaria",
	38:  "Barbados",
	39:  "China Highest",
	40:  "G99",
	41:  "Sweden",
	42:  "Chile",
	43:  "Brazil LV",
	44:  "NewZealand",
	45:  "IEEE1547 208VAC",
	46:  "IEEE1547 220VAC",
	47:  "IEEE1547 240VAC",
	48:  "60Hz LV Default",
	49:  "50Hz LV Default",
	50:  "AU_WAPN",
	51:  "AU_MicroGrid",
	52:  "JP_50Hz",
	53:  "JP_60Hz",
	54:  "India Higher",
	55:  "DEWA LV",
	56:  "DEWA MV",
	57:  "Slovakia",
	58:  "GreenGrid",
	59:  "Hungary",
	60:  "Sri Lanka",
	61:  "Spain Islands",
	62:  "Ergon30K",
	63:  "Energex30K",
	64:  "IEEE1547 230/400V",
	65:  "IEC61727 60Hz",
	66:  "Switzerland",
	67:  "CEI-016",
	68:  "AU_Horizon",
	69:  "Cyprus",
	70:  "AU_SAPN",
	71:  "AU_Ausgrid",
	72:  "AU_Essential",
	73:  "AU_Pwcore&CitiPW",
	74:  "Hong Kong",
	75:  "Poland MV",
	76:  "Holland MV",
	77:  "Sweden MV",
	78:  "VDE4110",
	96:  "cUSA_208VacDefault",
	97:  "cUSA_240VacDefault",
	98:  "cUSA_208VacCA_SCE",
	99:  "cUSA_240VacCA_SCE",
	100: "cUSA_208VacCA_SDGE",
	101: "cUSA_240VacCA_SDGE",
	102: "cUSA_208VacCA_PGE",
	103: "cUSA_240VacCA_PGE",
	104: "cUSA_208VacHECO_14HO",
	105: "cUSA_240VacHECO_14HO0x69",
	106: "cUSA_208VacHECO_14HM",
	107: "cUSA_240VacHECO_14HM",
}

//goland:noinspection SpellCheckingInspection
var ErrorCodes = map[int]string{
	31: "Internal, Communication Failure",
	30: "EEPROM R/W Failure",
	29: "Fac Failure",
	28: "DSP communication failure",
	27: "PhaseAngleFailure",
	26: "",
	25: "Relay Check Failure",
	24: "",
	23: "Vac Consistency Failure",
	22: "Fac Consistency Failure",
	21: "",
	20: "Back-Up Over Load",
	19: "DC Injection High",
	18: "Isolation Failure",
	17: "Vac Failure",
	16: "External Fan Failure",
	15: "PV Over Voltage",
	14: "Utility Phase Failure",
	13: "Over Temperature",
	12: "InternalFan Failure",
	11: "DC Bus High",
	10: "Ground I Failure",
	9:  "Utility Loss",
	8:  "AC HCT Failure",
	7:  "Relay Device Failure",
	6:  "GFCI Device Failure",
	5:  "",
	4:  "GFCI Consistency Failure",
	3:  "DCI Consistency Failure",
	2:  "",
	1:  "AC HCT Check Failure",
	0:  "GFCI Device Check Failure",
}

var DiagStatusCodes = map[int]string{
	0:  "Battery voltage low",
	1:  "Battery SOC low",
	2:  "Battery SOC in back",
	3:  "BMS: Discharge disabled",
	4:  "Discharge time on",
	5:  "Charge time on",
	6:  "Discharge Driver On",
	7:  "BMS: Discharge current low",
	8:  "APP: Discharge current too low",
	9:  "Meter communication failure",
	10: "Meter connection reversed",
	11: "Self-use load light",
	12: "EMS: discharge current is zero",
	13: "Discharge BUS high PV voltage",
	14: "Battery Disconnected",
	15: "Battery Overcharged",
	16: "BMS: Temperature too high",
	17: "BMS: Charge too high",
	18: "BMS: Charge disabled",
	19: "Self-use off",
	20: "SOC delta too volatile",
	21: "Battery self discharge too high",
	22: "Battery SOC low (off-grid)",
	23: "Grid wave unstable",
	24: "Export power limit set",
	25: "PF value set",
	26: "Real power limit set",
	27: "DC output on",
	28: "SOC protect off",
}

//goland:noinspection SpellCheckingInspection
var BmsAlarmCodes = map[int]string{
	15: "Charging over-voltage 3",
	14: "Discharging under-voltage 3",
	13: "Cell temperature high 3",
	12: "Communication failure 2",
	11: "Charging circuit failure",
	10: "Discharging circuit failure",
	9:  "Battery lock",
	8:  "Battery break",
	7:  "DC bus fault",
	6:  "Precharge fault",
	5:  "Discharging over-current 2",
	4:  "Charging over-current 2",
	3:  "Cell temperature low 2",
	2:  "Cell temperature high 2",
	1:  "Discharging under-voltage 2",
	0:  "Charging over-voltage 2",
}

var BmsWarningCodes = map[int]string{
	11: "System temperature high",
	10: "System temperature low 2",
	9:  "System temperature low 1",
	8:  "Cell imbalance",
	7:  "System reboot",
	6:  "Communication failure 1",
	5:  "Discharging over-current 1",
	4:  "Charging over-current 1",
	3:  "Cell temperature low 1",
	2:  "Cell temperature high 1",
	1:  "Discharging under-voltage 1",
	0:  "Charging over-voltage 1",
}
