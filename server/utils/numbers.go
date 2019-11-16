package utils

import (
	"math"
)

var eps float64 = 0.00000001

func FloatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}
