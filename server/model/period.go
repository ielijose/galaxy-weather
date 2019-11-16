package model

type Period struct {
	Start       uint   `json:"start"`
	End         uint   `json:"end"`
	Weather     string `json:"weather"`
	WeatherType Type   `json:"-"`
	Peak        *uint  `json:"peak,omitempty"`
}

type Stats struct {
	Drought   uint   `json:"drought"`
	Optimal   uint   `json:"optimal"`
	Rain      uint   `json:"rain"`
	RainPeaks []uint `json:"rain_peaks"`
}
