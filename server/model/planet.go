package model

import (
	"galaxy-weather/config"
	"galaxy-weather/utils"
	"math"
)

type direction uint8

const (
	Clockwise direction = iota
	CounterClockwise
)

type Planet struct {
	ID              uint      `gorm:"primary_key"`
	Name            string    `json:"name"`
	AngularVelocity uint      `json:"angular_velocity"`
	Distance        uint      `json:"distance"`
	Direction       direction `json:"direction"`
	Radio           int       `json:"radio"`
}

func (p Planet) AngularPosition(day uint) uint {
	degree := (day * p.AngularVelocity) % config.CircumferenceDegrees

	if p.Direction == Clockwise && degree > 0 {
		degree = config.CircumferenceDegrees - degree
	}

	return degree
}

func (p Planet) GetPointByDay(day uint) Position {
	degree := p.AngularPosition(day)

	radians := utils.DegreeToRadian(degree)

	x := math.Cos(radians) * float64(p.Distance)
	y := math.Sin(radians) * float64(p.Distance)

	return Position{
		PlanetID: p.ID,
		Day:      day,
		X:        x,
		Y:        y,
	}
}

func (p Planet) GetDaysPerYear() uint {
	return config.CircumferenceDegrees / p.AngularVelocity
}
