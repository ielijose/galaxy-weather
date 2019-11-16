package model

import (
	"math"
)

type Position struct {
	Day      uint    `json:"day" gorm:"default:0;primary_key;auto_increment:false"`
	PlanetID uint    `json:"planet_id" gorm:"primary_key;auto_increment:false"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

func (p Position) DistanceOf(p2 Position) float64 {
	return math.Sqrt(math.Pow(math.Abs(p2.X-p.X), 2) + math.Pow(math.Abs(p2.Y-p.Y), 2))
}
