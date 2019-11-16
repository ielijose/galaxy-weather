package utils

import (
	"math"
)

func DegreeToRadian(angle uint) float64 {
	return (float64(angle) * math.Pi) / 180
}
