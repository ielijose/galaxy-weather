package model

type Weather struct {
	Day         uint       `json:"day" gorm:"default:0;primary_key;auto_increment:false" sql:"not null"`
	Weather     string     `json:"weather" sql:"not null"`
	WeatherType Type       `json:"-" sql:"-"`
	Positions   *[]Position `json:"positions,omitempty" sql:"-"`
}

func (w Weather) ToType() Type {
	switch w.Weather {
	case "Drought":
		return Drought
	case "Rain":
		return Rain
	case "Heavy Rain":
		return HeavyRain
	case "Optimal Temperature":
		return OptimalTemperature
	case "Unknown":
		return Unknown
	default:
		return Unknown
	}
}
