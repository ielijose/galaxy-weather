package model

type Type int

const (
	Drought Type = iota
	Rain
	HeavyRain
	OptimalTemperature
	Unknown
)

func (t Type) String() string {
	switch t {
	case Drought:
		return "Drought"
	case Rain:
		return "Rain"
	case HeavyRain:
		return "Heavy Rain"
	case OptimalTemperature:
		return "Optimal Temperature"
	case Unknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}
